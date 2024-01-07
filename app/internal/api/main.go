package api

import (
	"html/template"
	"net/http"

	"example.com/internal/data"
)

func SwmsSchemaHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	var err error
	const error_msg = "<h2>Internal Server Error</h2>"

	tmpl, err = template.ParseFiles("views/components/swms_schema.html")

	if err != nil {
		// Log the error and return
		http.Error(w, error_msg, 500)
		return
	}

	tmp_err := tmpl.Execute(w, data.SwmsSchema)

	if tmp_err != nil {
		// Log the error
		http.Error(w, error_msg, 500)
	}

}
