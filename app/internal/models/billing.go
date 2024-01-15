package models

type SubscriptionPlans struct {
	ID                int    `json:"id"`
	StripePlanId      string `json:"stripe_plan_id"`
	StripePaymentLink string `json:"stripe_payment_link"`
	Description       string `json:"description"`
	Price             int    `json:"price"`
	Duration          string `json:"duration"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type Subscriptions struct {
	ID                   int     `json:"id"`
	OrganisationID       int     `json:"organisation_id"`
	SubscriptionPlanID   int     `json:"subscription_plan_id"`
	StripeSubscriptionID string  `json:"stripe_subscription_id"`
	StripeStatus         string  `json:"stripe_status"`
	CurrentPeriodStart   string  `json:"current_period_start"`
	CurrentPeriodEnd     string  `json:"current_period_end"`
	CancelledAt          *string `json:"cancelled_at"`
}

type StripeInvoices struct {
	ID              int    `json:"id"`
	SubscriptionID  int    `json:"subscription_id"`
	StripeInvoiceID string `json:"stripe_invoice_id"`
	StripeStatus    string `json:"stripe_status"`
	StripePDFLink   string `json:"stripe_pdf_link"`
}
