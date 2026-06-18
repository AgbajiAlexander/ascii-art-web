package handlers

import (
	"html/template"
	"net/http"

	"ascii-art-web/models"
)

var Templates *template.Template

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	data := models.PageData{
		Banner: "standard",
	}

	if err := Templates.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
	}
}
