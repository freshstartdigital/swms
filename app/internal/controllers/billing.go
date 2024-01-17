package controllers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"example.com/internal/models"
	"example.com/internal/repository"
	"example.com/internal/services"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	var err error

	tmpl, err = template.ParseFiles("views/register.html")
	if err != nil {
		// Log the error and return
		http.Error(w, "Internal Server Error", 500)
		return
	}
	tmpl.Execute(w, nil) // Execute with nil or appropriate context

}

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	var err error

	user, err := services.Authenticate(r)

	if err != nil {
		tmpl, err = template.ParseFiles("views/unauthenticated.html")
		if err != nil {
			// Log the error and return
			http.Error(w, "Internal Server Error", 500)
			return
		}
		tmpl.Execute(w, nil) // Execute with nil or appropriate context
		return
	}

	db := repository.ConnectDB()

	defer db.Close()

	organisation, err := db.GetOrganisation(user.OrganisationID)

	if err != nil {
		// Log the error and return
		http.Error(w, "Internal Server Error", 500)
		return
	}

	type AccountPageResponse struct {
		User               models.Users
		Organisation       models.Organisations
		CurrentPeriodEnd   string
		CurrentPeriodStart string
	}

	var currentPeriodStart string
	var currentPeriodEnd string

	subscriptions, errorMessage := db.GetSubscriptionByOrgID(organisation.ID)

	if errorMessage != nil {
		log.Println("err fetching sub", errorMessage)
		currentPeriodEnd = "N/A"
		currentPeriodStart = "N/A"
	} else {

		currentPeriodStart = time.Unix(subscriptions.CurrentPeriodStart, 0).Format("02/01/2006")
		currentPeriodEnd = time.Unix(subscriptions.CurrentPeriodEnd, 0).Format("02/01/2006")
	}

	accountPageResponse := AccountPageResponse{
		User:               user,
		Organisation:       organisation,
		CurrentPeriodEnd:   currentPeriodEnd,
		CurrentPeriodStart: currentPeriodStart,
	}

	tmpl, err = template.ParseFiles("views/account.html")
	if err != nil {
		// Log the error and return
		http.Error(w, "Internal Server Error", 500)
		return
	}
	tmpl.Execute(w, accountPageResponse) // Execute with nil or appropriate context

}
