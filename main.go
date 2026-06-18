package main

import (
	"html/template"
	"log"
	"net/http"

	"ascii-art-web/handlers"
)

func main() {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	handlers.InitTemplates(tmpl)

	http.HandleFunc("/", handlers.HandleIndex)
	http.HandleFunc("/ascii-art", handlers.HandleAsciiArt)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
