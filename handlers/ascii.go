package handlers

import (
	"net/http"

	"ascii-art-web/models"
	"ascii-art-web/services"
)

func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		render(w, http.StatusBadRequest, "index.html", models.PageData{
			Banner: banner,
			Error:  "empty input",
		})
		return
	}

	valid := map[string]struct{}{
		"standard":   {},
		"shadow":     {},
		"thinkertoy": {},
	}

	if _, ok := valid[banner]; !ok {
		render(w, http.StatusBadRequest, "index.html", models.PageData{
			InputText: text,
			Error:     "invalid banner",
		})
		return
	}

	chars, err := services.LoadBanner("banners/" + banner + ".txt")
	if err != nil {
		render(w, http.StatusNotFound, "index.html", models.PageData{
			InputText: text,
			Error:     err.Error(),
		})
		return
	}

	result, err := services.GenerateAsciiArt(text, chars)
	if err != nil {
		render(w, http.StatusBadRequest, "index.html", models.PageData{
			InputText: text,
			Error:     err.Error(),
		})
		return
	}

	render(w, http.StatusOK, "index.html", models.PageData{
		InputText: text,
		Banner:    banner,
		Result:    result,
	})
}
