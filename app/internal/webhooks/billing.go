package webhooks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"example.com/internal/models"
	"example.com/internal/repository"
	"example.com/internal/services"
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
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error getting organisation: %v\n", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

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
	case "product.created":
		type ProductCreated struct {
			ID          string  `json:"id"`
			Name        string  `json:"name"`
			Description *string `json:"description"`
		}

		var productCreated ProductCreated
		err := json.Unmarshal(stripe_data, &productCreated)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var description string
		if productCreated.Description != nil {
			description = *productCreated.Description
		} else {
			description = ""
		}

		err = db.CreateProduct(
			productCreated.ID,
			productCreated.Name,
			description)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating product: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case "price.created":
		type PriceCreated struct {
			ProductID  string `json:"product"`
			UnitAmount int    `json:"unit_amount"`
		}

		var priceCreated PriceCreated
		err := json.Unmarshal(stripe_data, &priceCreated)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.UpdateSubscriptionPricing(
			priceCreated.ProductID,
			priceCreated.UnitAmount,
		)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating subscription pricing: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case "price.updated":
		type PriceCreated struct {
			ProductID  string `json:"product"`
			UnitAmount int    `json:"unit_amount"`
		}

		var priceCreated PriceCreated
		err := json.Unmarshal(stripe_data, &priceCreated)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.UpdateSubscriptionPricing(
			priceCreated.ProductID,
			priceCreated.UnitAmount,
		)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating subscription pricing: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case "payment_link.created":
		type PaymentLinkCreated struct {
			ID  string `json:"id"`
			URL string `json:"url"`
		}
		log.Println("Payment link created")
		var paymentLinkCreated PaymentLinkCreated
		err := json.Unmarshal(stripe_data, &paymentLinkCreated)
		if err != nil {
			log.Println("Error parsing webhook JSON:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ProductID, err := services.GetProductIDFromPaymentLink(paymentLinkCreated.ID)

		if err != nil {
			log.Println("Error getting product ID from payment link:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.UpdateStripePaymentLink(
			ProductID,
			paymentLinkCreated.URL,
		)

		if err != nil {
			log.Println("Error updating stripe payment link:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	w.WriteHeader(http.StatusOK)

}
