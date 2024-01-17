package models

type SubscriptionPlans struct {
	ID                int     `json:"id"`
	StripeProductId   string  `json:"stripe_product_id"`
	StripePaymentLink *string `json:"stripe_payment_link"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             int     `json:"price"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}

type Subscriptions struct {
	ID                   int     `json:"id"`
	OrganisationID       int     `json:"organisation_id"`
	SubscriptionPlanID   int     `json:"subscription_plan_id"`
	StripeSubscriptionID string  `json:"stripe_subscription_id"`
	StripeStatus         string  `json:"stripe_status"`
	CurrentPeriodStart   int64   `json:"current_period_start"`
	CurrentPeriodEnd     int64   `json:"current_period_end"`
	CancelledAt          *string `json:"cancelled_at"`
}
