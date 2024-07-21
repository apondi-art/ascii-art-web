package handlers

import (
	"html/template"
	"log"
	"net/http"

	art "ascii-art-web/ascii-art"
)

type Data struct {
	Filename string
	Input    string
	Result   string
}

func PostAsciiArt(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request on /ascii-art route\n", r.Method)

	if r.Method != http.MethodPost {
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	action := r.FormValue("action")
	if action == "Reset" {
		// Clear the form data by rendering the template with empty data
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

	text := r.FormValue("input")
	banner := r.FormValue("filename")

	if text == "" || banner == "" {
		http.Error(w, "400 Bad Request: Missing text or banner", http.StatusBadRequest)
		return
	}

	result, err := art.AsciiArt(text, banner)
	if err != nil {
		log.Printf("Error generating ASCII art: %v\n", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

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
