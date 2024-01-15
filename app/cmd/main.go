package main

import (
	"net/http"

	"example.com/internal/api"
	"example.com/internal/controllers"
	"example.com/internal/webhooks"
	"github.com/gorilla/mux"
)

// HomePageData represents data to be passed to the template
type HomePageData struct {
	Message string
}

func main() {
	r := mux.NewRouter()

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Route controllers
	r.HandleFunc("/", controllers.HomePageHandler).Methods("GET")
	r.HandleFunc("/create", controllers.CreatePageHandler).Methods("GET")
	r.HandleFunc("/login", controllers.GetLoginHandler).Methods("GET")
	r.HandleFunc("/register", controllers.RegisterHandler).Methods("GET")

	// routes for API
	r.HandleFunc("/api/swms/schema", api.SwmsSchemaHandler).Methods("GET")
	r.HandleFunc("/api/swms", api.GetSwms).Methods("GET")
	r.HandleFunc("/api/swms", api.CreateSwms).Methods("POST")
	r.HandleFunc("/api/swms", api.UpdateFileHandler).Methods("PATCH")
	r.HandleFunc("/api/login", api.LoginHandler).Methods("POST")
	r.HandleFunc("/api/register", api.RegisterOrgHandler).Methods("POST")

	// Webhooks
	r.HandleFunc("/webhooks/billing", webhooks.BillingWebhookHandler).Methods("POST")
	// Listen and Serve using the mux router
	http.ListenAndServe(":8080", r)
}
