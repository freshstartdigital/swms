package api

import (
	"html/template"
	"log"
	"net/http"

	"example.com/internal/models"
	"example.com/internal/repository"
	"example.com/internal/services"
)

func GetSwms(w http.ResponseWriter, r *http.Request) {
	var templ *template.Template
	var swms []models.Swms
	var err error

	db := repository.ConnectDB()

	defer db.Close()

	user, err := services.Authenticate(r)

	swms, err = db.GetAllSwms(user.OrganisationID)

	if err != nil {
		// Log the error and return
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error getting all swms: %v", err)
		return
	}

	templ, err = template.ParseFiles("views/components/home/table.html")

	if err != nil {
		// Log the error and return
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}

	for i := 0; i < len(swms); i++ {
		services.GetFilePath(&swms[i])
	}

	err = templ.Execute(w, swms)

	if err != nil {
		// Log the error and return
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
		return
	}

	return

}
