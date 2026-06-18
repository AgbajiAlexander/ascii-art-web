package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// PageData holds everything we pass to the HTML template
type PageData struct {
	InputText string // what the user typed
	Banner    string // which banner they chose
	Result    string // the ASCII art result
	Error     string // any error message
}

// templates — loaded once at startup
var templates *template.Template

func main() {
	// Load all HTML templates from the templates folder
	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Could not load templates:", err)
	}

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/ascii-art", handleAsciiArt)

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Reject any path that isn't exactly "/"
	if r.URL.Path != "/" {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Render the template with default values
	data := PageData{Banner: "standard"}
	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data from the POST request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "400 - Bad Request", http.StatusBadRequest)
		return
	}

	// Get the form values
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// Validate inputs
	if text == "" {
		data := PageData{
			Banner: banner,
			Error:  "Please enter some text",
		}
		w.WriteHeader(http.StatusBadRequest)
		templates.ExecuteTemplate(w, "index.html", data)
		return
	}

	// Validate banner choice
	validBanners := map[string]bool{
		"standard": true, "shadow": true, "thinkertoy": true,
	}
	if !validBanners[banner] {
		data := PageData{
			InputText: text,
			Error:     "Invalid banner selection",
		}
		w.WriteHeader(http.StatusBadRequest)
		templates.ExecuteTemplate(w, "index.html", data)
		return
	}

	// Load the banner file
	chars, err := loadBanner(banner + ".txt")
	if err != nil {
		data := PageData{
			InputText: text,
			Banner:    banner,
			Error:     "Banner file not found",
		}
		w.WriteHeader(http.StatusNotFound)
		templates.ExecuteTemplate(w, "index.html", data)
		return
	}

	// Generate ASCII art
	result := generateAsciiArt(text, chars)

	// Send result back to page
	data := PageData{
		InputText: text,
		Banner:    banner,
		Result:    result,
	}
	w.WriteHeader(http.StatusOK)
	templates.ExecuteTemplate(w, "index.html", data)
}

func loadBanner(filename string) ([][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("banner file is empty")
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	var chars [][]string
	for i := 0; i+8 <= len(lines); i += 9 {
		chars = append(chars, lines[i:i+8])
	}

	if len(chars) != 95 {
		return nil, fmt.Errorf("banner file is incomplete")
	}
	return chars, nil
}

func generateAsciiArt(input string, chars [][]string) string {
	// Handle \n in input
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	var sb strings.Builder
	for _, line := range lines {
		if line == "" {
			sb.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range line {
				index := int(ch) - 32
				if index >= 0 && index < len(chars) {
					sb.WriteString(chars[index][row])
				}
			}
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
