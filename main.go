package main

import (
	"log"
	"net/http"

	"github.com/neomede/porrify_api/handlers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Listening...")
	n := negroni.Classic()

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handlers.UserHandler).Methods("GET")
	r.HandleFunc("/users", handlers.AddUserHandler).Methods("POST")

	r.HandleFunc("/circuits", handlers.AllCircuitsHandler).Methods("GET")
	r.HandleFunc("/circuits/{id}", handlers.CircuitHandler).Methods("GET")

	r.HandleFunc("/categories/{category}/pilots", handlers.PilotsHandler).Methods("GET")

	r.HandleFunc("/bets", handlers.AddBetHandler).Methods("POST")
	r.HandleFunc("/users/{user_id}/circuits/{circuit_id}/bets", handlers.BetHandler).Methods("GET")

	n.UseHandler(corsHandler(r))
	n.Run(":8888")
}

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		if r.Method == "OPTIONS" {
			//handle preflight in here
		} else {
			h.ServeHTTP(w, r)
		}
	}
}
