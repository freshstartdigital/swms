package main

import (
	"net/http"

	"example.com/internal/api"
	"example.com/internal/controllers"
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
	r.HandleFunc("/", controllers.HomePageHandler)
	r.HandleFunc("/create", controllers.CreatePageHandler)

	// routes for API
	r.HandleFunc("/api/swms/schema", api.SwmsSchemaHandler).Methods("POST")

	// Listen and Serve using the mux router
	http.ListenAndServe(":8080", r)
}
