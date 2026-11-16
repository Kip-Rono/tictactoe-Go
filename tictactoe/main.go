package main

import (
	"fmt"
	"net/http"
	"log"
)

func homeHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintln(
		w, "Welcome to the Tic Tac Toe Game",
	)

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