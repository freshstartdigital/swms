package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// PaymentLinkLineItem represents the structure of a line item in the payment link response
type PaymentLinkLineItem struct {
	ID             string `json:"id"`
	Object         string `json:"object"`
	AmountDiscount int    `json:"amount_discount"`
	AmountSubtotal int    `json:"amount_subtotal"`
	AmountTax      int    `json:"amount_tax"`
	AmountTotal    int    `json:"amount_total"`
	Currency       string `json:"currency"`
	Description    string `json:"description"`
	Price          struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"price"`
	Quantity int `json:"quantity"`
}

// PaymentLinkResponse represents the structure of the response from the payment link endpoint
type PaymentLinkResponse struct {
	Object  string                `json:"object"`
	Data    []PaymentLinkLineItem `json:"data"`
	HasMore bool                  `json:"has_more"`
	URL     string                `json:"url"`
}

// GetProductIDFromPaymentLink fetches the product ID associated with the provided payment link ID
func GetProductIDFromPaymentLink(paymentLinkID string) (string, error) {
	stripeSecretKey := os.Getenv("STRIPE_SECRET_KEY")
	log.Println("Stripe secret key:", stripeSecretKey)
	// Construct the URL for the Stripe endpoint
	stripeURL := fmt.Sprintf("https://api.stripe.com/v1/payment_links/%s/line_items", paymentLinkID)

	// Create an HTTP client with the necessary headers
	client := &http.Client{}
	req, err := http.NewRequest("GET", stripeURL, nil)
	if err != nil {
		return "", err
	}

	// Set the Authorization header with the Stripe secret key
	req.Header.Set("Authorization", "Bearer "+stripeSecretKey)

	// Make the GET request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the response into a PaymentLinkResponse object
	var paymentLinkResponse PaymentLinkResponse
	err = json.Unmarshal(body, &paymentLinkResponse)
	if err != nil {
		return "", err
	}

	// Check if there is at least one line item in the response
	if len(paymentLinkResponse.Data) > 0 {
		// Retrieve the Product ID from the first line item
		productID := paymentLinkResponse.Data[0].Price.ID
		return productID, nil
	}

	return "", fmt.Errorf("No line items found in the payment link response")
}
