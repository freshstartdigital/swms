package models

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Session struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	SessionToken string `json:"session_token"`
	ExpiresAt    string `json:"expires_at"`
}
