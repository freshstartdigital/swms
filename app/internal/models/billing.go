package models

type KeyStatus string

const (
	Active  KeyStatus = "active"
	Expired KeyStatus = "expired"
	Revoked KeyStatus = "revoked"
)

type LicenceKeys struct {
	ID             int       `json:"id"`
	OrganisationID *int      `json:"organisation_id"`
	LicenceKey     string    `json:"licence_key"`
	KeyStatus      KeyStatus `json:"key_status"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
}

type SubscriptionPlans struct {
	ID           int    `json:"id"`
	StripePlanID string `json:"stripe_plan_id"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	Duration     string `json:"duration"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
