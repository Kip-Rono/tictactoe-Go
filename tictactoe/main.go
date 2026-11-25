package main

import (
	
	"log"
	"net/http"
	"text/template"
)

func homeHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	tmpl, err := template.ParseFiles("templates/home.html")

	if err != nil {
		http.Error(
			w,
			"Template Error",
			http.StatusInternalServerError,
		)
		return
	}
		
	

	err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(
			w,
			"unable to render page",
			http.StatusInternalServerError,
		)
	}

}

func main() {

	http.HandleFunc("/", homeHandler)

	log.Println("App running on port 8080")

	log.Fatal(
		http.ListenAndServe(
			":8080",
			nil,
		),
	)

}