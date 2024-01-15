package api

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"example.com/internal/models"
	"example.com/internal/repository"
	"example.com/internal/services"
)

func RegisterOrgHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	var err error

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	businessName := r.FormValue("businessName")

	db := repository.ConnectDB()

	defer db.Close()

	var organisationID int
	var userID int

	organisationID, err = db.CreateOrganisation(businessName)

	if err != nil {
		tmpl, err = template.ParseFiles("views/components/register/error.html")
		if err != nil {
			// Log the error and return
			http.Error(w, "Internal Server Error", 500)
			return
		}
		tmpl.Execute(w, "Error creating organisation") // Execute with nil or appropriate context
		return
	}

	userID, err = db.CreateUser(name, email, password, organisationID)

	if err != nil {
		tmpl, err = template.ParseFiles("views/components/register/error.html")
		if err != nil {
			// Log the error and return
			http.Error(w, "Internal Server Error", 500)
			return
		}

		tmpl.Execute(w, "Error creating user") // Execute with nil or appropriate context
		return
	}

	err = db.UpdateAccountHolderID(userID, organisationID)

	if err != nil {
		tmpl, err = template.ParseFiles("views/components/register/error.html")
		if err != nil {
			// Log the error and return
			http.Error(w, "Internal Server Error", 500)
			return
		}
		tmpl.Execute(w, "Error creating account") // Execute with nil or appropriate context
		return
	}

	var subscriptionPlans models.SubscriptionPlans

	subscriptionPlans, err = db.GetSubscriptionPlanByID(1)

	if err != nil {
		tmpl, err = template.ParseFiles("views/components/register/error.html")
		if err != nil {
			// Log the error and return
			http.Error(w, "Internal Server Error", 500)
			return
		}
		log.Printf("Error getting subscription plan: %v", err)
		tmpl.Execute(w, "Error fetching data") // Execute with nil or appropriate context
		return
	}

	tmpl, err = template.ParseFiles("views/components/register/success.html")
	if err != nil {
		// Log the error and return
		http.Error(w, "Internal Server Error", 500)
		return
	}

	var sessionToken string
	sessionToken, err = services.GenerateSessionToken()
	err = db.CreateSession(sessionToken, userID, r.RemoteAddr)

	http.SetCookie(w, &http.Cookie{
		Name:    "swms_session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/", // Global path
	})

	type RegisterOrgPageResponse struct {
		SubscriptionPlans models.SubscriptionPlans
		OrganisationID    int
	}
	tmplRes := RegisterOrgPageResponse{
		SubscriptionPlans: subscriptionPlans,
		OrganisationID:    organisationID,
	}

	tmpl.Execute(w, tmplRes) // Execute with nil or appropriate context

}
