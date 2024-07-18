package main

import (
	"fmt"
	"net/http"
	"text/template"

	handle "ascii-art-web/handler"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", Htmlhandler)
	http.HandleFunc("/ascii-art", handle.Handle)
	fmt.Println("server created successfully")
	http.ListenAndServe(":8080", nil)
}

func Htmlhandler(w http.ResponseWriter, r *http.Request) {
	var getv handle.Outpt
	v, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Fprint(w, err)
	}
	v.Execute(w, getv)
}
