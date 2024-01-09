package controllers

import (
	"html/template"
	"log"
	"net/http"

	"example.com/internal/data"
	"example.com/internal/models"
	"example.com/internal/repository"
	"example.com/internal/services"
	"github.com/gorilla/mux"
)

type HomePageResponse struct {
	Swms []models.Swms
}

// HomePageHandler handles the home page requests
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	_, err := services.Authenticate(r)

	if err != nil {
		tmpl, err = template.ParseFiles("views/unauthenticated.html")
		if err != nil {
			// Log the error and return
			log.Printf("Error parsing template: %v", err)
			http.Error(w, "Internal Server Error", 500)
			return
		}
		tmpl.Execute(w, nil)
		return
	}
	tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		// Log the error and return
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	db := repository.ConnectDB()
	defer db.Close()
	swms, err := db.GetAllSwms()
	if err != nil {
		log.Printf("Error getting all swms: %v", err)

		swms = []models.Swms{}
	}

	for i := 0; i < len(swms); i++ {
		services.GetFilePath(&swms[i])
	}

	data := HomePageResponse{
		Swms: swms,
	}

	tmp_err := tmpl.Execute(w, data)
	if tmp_err != nil {
		// Log the error
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func CreatePageHandler(w http.ResponseWriter, r *http.Request) {
	sitename := mux.Vars(r)["sitename"]
	var tmpl *template.Template
	var err error
	if sitename == "" {
		tmpl, err = template.ParseFiles("views/create.html")
		// Log the error and return
		if err != nil {
			log.Printf("Error parsing template: %v", err)
			http.Error(w, "Internal Server Error", 500)
			return
		}

	}

	data := data.SwmsSchema

	tmp_err := tmpl.Execute(w, data)
	if tmp_err != nil {
		// Log the error
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	var err error
	tmpl, err = template.ParseFiles("views/login.html")
	// Log the error and return
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	tmp_err := tmpl.Execute(w, nil)
	if tmp_err != nil {
		// Log the error
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	_, err := services.Login(email, password, w)
	if err != nil {
		log.Printf("Error logging in: %v", err)
		http.Error(w, "Internal Server Error", 500)
	}
	http.Redirect(w, r, "/", 302)
}
