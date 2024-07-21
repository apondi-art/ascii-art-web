// Package main is the entry point of this ASCII art web application.
// It contains all the functions that handle forms, CSS, and ASCII art generation.
// It also creates the server to listen and serve.

package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("static")) //Serve static files from the "static" directory at the "/static/" URL path.
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.GetAsciiForm)          //Handle the root URL path("/") to display the ASCII art form.
	http.HandleFunc("/ascii-art", handlers.PostAsciiArt) //Handle the "/ascii-art"URL path to process form submission.
	fmt.Println("SUCCESS!! listen to server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil { //start the server and listenon port 8080.
		log.Fatal("Server failed to start:", err)
	}
}
