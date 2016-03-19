package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neomede/porrify_api/handlers"
)

func main() {
	fmt.Println("Listening in 8080...")

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handlers.UserHandler).Methods("GET")
	r.HandleFunc("/users", handlers.AddUserHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}
