package handlers

import (
	"net/http"

	"ascii-art-web/models"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := models.PageData{
		Banner: "standard",
	}

	render(w, http.StatusOK, "index.html", data)
}
