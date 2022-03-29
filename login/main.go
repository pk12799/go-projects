package main

import (
	"fmt"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	// w.Header("content-type", "application/json")
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Fprintf(w, "email = %s\n", email)
	fmt.Fprintf(w, "password = %s\n", password)
}

func main() {
	file := http.FileServer(http.Dir("./template"))
	http.Handle("/", file)
	http.HandleFunc("/login", login)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
