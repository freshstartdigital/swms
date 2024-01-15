package models

type OrganisationTypes struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	DisplayName string `json:"display_name"`
}

type Organisations struct {
	ID               int                `json:"id"`
	OrganisationType *OrganisationTypes `json:"organisation_type"`
	Name             string             `json:"name"`
	BusinessAddress  string             `json:"business_address"`
	ABN              string             `json:"abn"`
	BusinessPhone    string             `json:"business_phone"`
	BusinessEmail    string             `json:"business_email"`
	LogoFileName     string             `json:"logo_file_name"`
	CreatedAt        string             `json:"created_at"`
	UpdatedAt        string             `json:"updated_at"`
	AccountHolderID  *int               `json:"account_holder_id"`
	StripeCustomerID string             `json:"stripe_customer_id"`
	StripeUrl        string             `json:"stripe_url"`
}

type Users struct {
	ID             int    `json:"id"`
	OrganisationID int    `json:"organisation_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type Session struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	SessionToken string `json:"session_token"`
	IPAddress    string `json:"ip_address"`
	CreatedAt    string `json:"created_at"`
	ExpiresAt    string `json:"expires_at"`
}

type LoginAttempts struct {
	ID        int    `json:"id"`
	IPAddress string `json:"ip_address"`
	CreatedAt string `json:"created_at"`
}

type BannedIPs struct {
	ID        int    `json:"id"`
	IPAddress string `json:"ip_address"`
	CreatedAt string `json:"created_at"`
	ExpiresAt string `json:"expires_at"`
}

type PasswordResets struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
}
