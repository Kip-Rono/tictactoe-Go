package main

import (
	
	"log"
	"net/http"
	"text/template"
)

var (
	homeTemplate *template.Template
	loginTemplate *template.Template
)

func init () {
	// Parse layout, partials and pages together
	var err error

	homeTemplate, err = template.ParseFiles(
		"templates/layouts/app.html",
		"templates/partials/header.html",
		"templates/partials/footer.html",
		"templates/pages/home.html",
	)

	if err != nil {
		log.Fatal("Error parsing home template:", err)
	}

	loginTemplate, err = template.ParseFiles(
		"templates/layouts/app.html",
		"templates/partials/header.html",
		"templates/partials/footer.html",
		"templates/auth/login.html",
	)

	if err != nil {
		log.Fatal("Error parsing login template:", err)
	}
}

func main() {
	// Init database
	db, err := initDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer db.Close()

	// register static files
	fileServer := http.FileServer(
		http.Dir("static"),
	)

	// multiplexer to handle requests
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/login", loginHandler)

	log.Println("App running on port 8080")

	log.Fatal(
		http.ListenAndServe(":8080", mux),
	)
}

func homeHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	if err := homeTemplate.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println("Execute error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func loginHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	if err := loginTemplate.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println("Execute error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}