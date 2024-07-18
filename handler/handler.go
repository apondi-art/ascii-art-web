package handle

import (
	"fmt"
	"net/http"
	"text/template"

	asciiprint "ascii-art-web/ascii-art"
)

type Outpt struct {
	Filename string
	Input    string
	Result   string
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var values Outpt
	fil, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprint(w, err)
			return
		}

		values.Filename = r.FormValue("filename")
		values.Input = r.FormValue("input")
		values.Result, err = asciiprint.CheckPrint(values.Input, values.Filename)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

	}
	fil.Execute(w, values)
}
