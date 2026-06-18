package main

import (
	"fmt"
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

	handlers.Templates = tmpl

	http.HandleFunc("/", handlers.HandleIndex)
	http.HandleFunc("/ascii-art", handlers.HandleAsciiArt)

	fmt.Println("Server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
