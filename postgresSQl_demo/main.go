package main

import (
	"fmt"
	"net/http"
	"os"

	"v1/app"
	"v1/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // attach JWT auth middleware

	port := os.Getenv("PORT") // Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" // localhost
	}

	fmt.Println(port)
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")
	err := http.ListenAndServe(":"+port, router) // Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
