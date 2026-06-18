package handlers

import (
	"net/http"
)

func render(w http.ResponseWriter, status int, tmpl string, data any) {
	w.WriteHeader(status)

	if templates == nil {
		http.Error(w, "templates not initialized", http.StatusInternalServerError)
		return
	}

	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
	}
}
