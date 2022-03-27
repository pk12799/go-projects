package main

import (
	"net/http"
	"oauth/controller"
)

func main() {
	http.HandleFunc("/", controller.Home)

	http.HandleFunc("/google/login", controller.GoogleLogin)

	http.HandleFunc("/google/callback", controller.GoogleCallback)

	http.ListenAndServe(":3000", nil)
}
