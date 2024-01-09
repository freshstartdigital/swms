package services

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"example.com/internal/models"
	"example.com/internal/repository"
)

func generateSessionToken() (string, error) {
	b := make([]byte, 32) // 32 bytes generates a 256-bit random number
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func Login(email string, password string, w http.ResponseWriter) (models.User, error) {
	db := repository.ConnectDB()
	defer db.Close()

	user, err := db.LoginHandler(email, password)
	if err != nil {
		return user, err
	}

	// Generate a secure session token
	sessionToken, err := generateSessionToken()
	if err != nil {
		return user, err // Handle the error properly
	}

	// Store the session token in your session store with user details
	err = db.StoreSessionTokenHandler(sessionToken, user.ID)

	// Set the session token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "swms_session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(24 * time.Hour), // 1 day for example
		Path:    "/",                            // Global path
	})

	return user, nil
}

func Authenticate(r *http.Request) (models.User, error) {
	// Get the session cookie
	c, err := r.Cookie("swms_session_token")
	if err != nil {
		return models.User{}, err
	}

	// Get the session token value
	sessionToken := c.Value

	// Get the session token from your session store
	// If the session token is not present or invalid, return an error
	// Otherwise return the user ID stored with the session token
	db := repository.ConnectDB()
	defer db.Close()

	user, err := db.AuthenticateHandler(sessionToken)
	if err != nil {
		return user, err
	}

	return user, nil
}
