package api

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"example.com/internal/models"
	"example.com/internal/repository"
	"example.com/internal/services"
)

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	err = r.ParseForm()
	if err != nil {
		// Handle error
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	// Now use email and password

	var tmpl *template.Template

	var user models.Users

	db := repository.ConnectDB()

	defer db.Close()

	var banned models.BannedIPs
	var isBanned bool

	banned, isBanned, err = db.GetBannedIp(r.RemoteAddr)

	if err != nil {
		// Log the error and return
		log.Printf("Error getting banned IP: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if isBanned {
		tmpl, err = template.ParseFiles("views/components/login/banned.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
			http.Error(w, "Internal Server Error", 500)
			return
		}
		tmpl.Execute(w, banned) // Execute with nil or appropriate context
		return
	}

	user, err = db.GetUser(email, password)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			var attempts int
			attempts, err = db.CreateLoginAttempt(strings.Split(r.RemoteAddr, ":")[0])

			if err != nil {
				// Log the error and return
				log.Printf("Error creating login attempt: %v", err)
				http.Error(w, "Internal Server Error", 500)
				return
			}

			if attempts >= 5 {
				err = db.CreateBannedIp(strings.Split(r.RemoteAddr, ":")[0])
				tmpl, err = template.ParseFiles("views/components/login/banned.html")
				if err != nil {
					log.Printf("Error parsing template: %v", err)
					http.Error(w, "Internal Server Error", 500)
					return
				}
				tmpl.Execute(w, nil) // Execute with nil or appropriate context
				return
			}

			// Handle incorrect credentials
			tmpl, err = template.ParseFiles("views/components/login/invalid.html")
			if err != nil {
				log.Printf("Error parsing template: %v", err)
				http.Error(w, "Internal Server Error", 500)
				return
			}
			var attemptsRemaining int = 5 - attempts
			tmpl.Execute(w, attemptsRemaining) // Execute with nil or appropriate context
			return
		}

		// Handle other errors
		log.Printf("Error getting user: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	tmpl, err = template.ParseFiles("views/components/login/success.html")
	// Log the error and return
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	var sessionToken string
	sessionToken, err = services.GenerateSessionToken()

	if err != nil {
		// Log the error and return
		log.Printf("Error generating session token: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = db.CreateSession(sessionToken, user.ID, r.RemoteAddr)

	if err != nil {
		// Log the error and return
		log.Printf("Error creating session: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "swms_session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/", // Global path
	})

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusOK)

	tmp_err := tmpl.Execute(w, user)
	if tmp_err != nil {
		// Log the error
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", 500)
	}
}
