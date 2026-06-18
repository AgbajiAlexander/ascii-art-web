package handlers

import (
	"net/http"

	"ascii-art-web/models"
	"ascii-art-web/services"
)

func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "400 - Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		render(w, http.StatusBadRequest, models.PageData{
			Banner: banner,
			Error:  "Please enter text",
		})
		return
	}

	valid := map[string]struct{}{
		"standard":   {},
		"shadow":     {},
		"thinkertoy": {},
	}

	if _, ok := valid[banner]; !ok {
		render(w, http.StatusBadRequest, models.PageData{
			InputText: text,
			Error:     "Invalid banner",
		})
		return
	}

	chars, err := services.LoadBanner("banners/" + banner + ".txt")
	if err != nil {
		render(w, http.StatusNotFound, models.PageData{
			InputText: text,
			Banner:    banner,
			Error:     err.Error(),
		})
		return
	}

	result, err := services.GenerateAsciiArt(text, chars)
	if err != nil {
		render(w, http.StatusBadRequest, models.PageData{
			InputText: text,
			Banner:    banner,
			Error:     err.Error(),
		})
		return
	}

	render(w, http.StatusOK, models.PageData{
		InputText: text,
		Banner:    banner,
		Result:    result,
	})
}

func render(w http.ResponseWriter, status int, data models.PageData) {
	w.WriteHeader(status)
	Templates.ExecuteTemplate(w, "index.html", data)
}
