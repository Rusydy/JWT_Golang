package main

import (
	"log"

	"github.com/gorilla/mux"
)

// AddApproutes will add the routes for the application
func AddApproutes(route *mux.Router) {
	log.Println("Load Router...")

	route.HandleFunc("/signin", SignInUser).Methods("POST")
	route.HandleFunc("/signup", SignUpUser).Methods("POST")
	route.HandleFunc("/userDetails", GetUserDetails).Methods("GET")

	log.Println("Routes are Loaded.")
}
