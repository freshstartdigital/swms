package webhooks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"example.com/internal/models"
	"example.com/internal/repository"
	"github.com/stripe/stripe-go/v72"
)

func BillingWebhookHandler(w http.ResponseWriter, req *http.Request) {
	const MaxBodyBytes = int64(65536)
	req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	event := stripe.Event{}

	if err := json.Unmarshal(payload, &event); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse webhook body json: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stripe_data := event.Data.Raw

	db := repository.ConnectDB()
	defer db.Close()

	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case "customer.created":
		type CustomerCreated struct {
			ClientReferenceID *int   `json:"client_reference_id"`
			Email             string `json:"email"`
			ID                string `json:"id"`
			Subscriptions     struct {
				URL string `json:"url"`
			} `json:"subscriptions"`
		}

		var customerCreated CustomerCreated
		err := json.Unmarshal(stripe_data, &customerCreated)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var organisation models.Organisations

		if customerCreated.ClientReferenceID != nil {
			organisation, err = db.GetOrganisation(*customerCreated.ClientReferenceID)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error getting organisation: %v\n", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		} else {
			organisation, err = db.GetOrganisationByUserEmail(customerCreated.Email)
		}

		if organisation.StripeCustomerID != "" {
			fmt.Fprintf(os.Stderr, "Organisation already has a Stripe Customer ID: %v\n", organisation.StripeCustomerID)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.UpdateOrganisationStripeData(organisation.ID, customerCreated.ID, customerCreated.Subscriptions.URL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating organisation: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case "customer.subscription.created":

		type CustomerSubscriptionCreated struct {
			CustomerID         string `json:"customer"`
			Status             string `json:"status"`
			SubscriptionID     string `json:"id"`
			CurrentPeriodStart int64  `json:"current_period_start"`
			CurrentPeriodEnd   int64  `json:"current_period_end"`
			Items              struct {
				Data []struct {
					Plan struct {
						ProductID string `json:"product"`
					} `json:"plan"`
				} `json:"data"`
			} `json:"items"`
		}

		var customerSubscriptionCreated CustomerSubscriptionCreated
		err := json.Unmarshal(stripe_data, &customerSubscriptionCreated)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.CreateSubscription(
			customerSubscriptionCreated.CustomerID,
			customerSubscriptionCreated.Status,
			customerSubscriptionCreated.SubscriptionID,
			customerSubscriptionCreated.Items.Data[0].Plan.ProductID,
			customerSubscriptionCreated.CurrentPeriodStart,
			customerSubscriptionCreated.CurrentPeriodEnd,
		)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating subscription: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case "customer.subscription.deleted":
		type CustomerSubscriptionDeleted struct {
			ID          string `json:"id"`
			CancelledAt int64  `json:"cancel_at"`
		}

		var customerSubscriptionDeleted CustomerSubscriptionDeleted
		err := json.Unmarshal(stripe_data, &customerSubscriptionDeleted)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.CreateCancelledAtSubscription(
			customerSubscriptionDeleted.ID,
			customerSubscriptionDeleted.CancelledAt,
		)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating cancelled_at subscription: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case "invoice.created":
		type InvoiceCreated struct {
			SubscriptionID string `json:"subscription"`
			Status         string `json:"status"`
			ID             string `json:"id"`
			InvoicePDF     string `json:"invoice_pdf"`
		}

		var invoiceCreated InvoiceCreated
		err := json.Unmarshal(stripe_data, &invoiceCreated)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.CreateInvoice(
			invoiceCreated.SubscriptionID,
			invoiceCreated.ID,
			invoiceCreated.Status,
			invoiceCreated.InvoicePDF,
		)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating invoice: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case "invoice.paid":
		type InvoicePaid struct {
			Status string `json:"status"`
			ID     string `json:"id"`
			Lines  struct {
				Data []struct {
					Period struct {
						Start int64 `json:"start"`
						End   int64 `json:"end"`
					} `json:"period"`
				} `json:"data"`
			} `json:"lines"`
		}

		var invoicePaid InvoicePaid
		err := json.Unmarshal(stripe_data, &invoicePaid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.UpdateInvoice(
			invoicePaid.ID,
			invoicePaid.Status,
		)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating invoice: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.UpdateSubscription(
			invoicePaid.ID,
			invoicePaid.Status,
			invoicePaid.Lines.Data[0].Period.Start,
			invoicePaid.Lines.Data[0].Period.End,
		)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating subscription: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	w.WriteHeader(http.StatusOK)

}
