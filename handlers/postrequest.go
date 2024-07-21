// Package handlers contains the function for handling POST requests to generate ASCII art.
package handlers

import (
	"html/template"
	"log"
	"net/http"

	art "ascii-art-web/ascii-art"
)

// Data represents the data structure used for rendering the form template.
type Data struct {
	Filename string
	Input    string
	Result   string
}

// PostAsciiArt handles POST requests to the "/ascii-art" route.
// It processes the form submission to generate ASCII art or reset the form.
func PostAsciiArt(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request on /ascii-art route\n", r.Method)

	if r.Method != http.MethodPost {
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	action := r.FormValue("action")
	if action == "Reset" {
		// Clear the form data by rendering the template with empty data
		// Handle form actions, such as "Reset" to clear the form.
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Printf("Error parsing template: %v\n", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, &Data{})
		if err != nil {
			log.Printf("Error executing template: %v\n", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}
	// Retrieve input values from the form submission.
	text := r.FormValue("input")
	banner := r.FormValue("filename")

	if text == "" || banner == "" {
		http.Error(w, "400 Bad Request: Missing text or banner", http.StatusBadRequest)
		return
	}
	// Generate ASCII art using the provided input and banner filename.
	result, err := art.AsciiArt(text, banner)
	if err != nil {
		log.Printf("Error generating ASCII art: %v\n", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Prepare the data for rendering the template with results.
	resultData := &Data{
		Filename: banner,
		Input:    text,
		Result:   result,
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, resultData)
	if err != nil {
		log.Printf("Error executing template: %v\n", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}
