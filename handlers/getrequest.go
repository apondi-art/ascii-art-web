// Package handlers contains the function for handling GET requests for the ASCII form submission.
package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// GetAsciiForm handles GET requests for the root path ("/") by rendering the ASCII form template.
// It returns an error response for non-GET requests and other paths.
func GetAsciiForm(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		// Only allow GET requests; respond with a 405 error for other methods.
		if r.Method != http.MethodGet {
			http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		// Attempt to parse the template file.
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			// Log parsing errors and return a 500 Internal Server Error.
			log.Printf("Error parsing template: %v\n", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Execute the parsed template and write it to the response.
		err = t.Execute(w, nil)
		if err != nil {
			// Log execution errors and return a 500 Internal Server Error.
			log.Printf("Error executing template: %v\n", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

	default:
		// Return a 404 Not Found error for all other paths.
		http.NotFound(w, r)
	}
}
