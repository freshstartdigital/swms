package controllers

import (
	"html/template"
	"net/http"
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
