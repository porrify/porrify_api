package porrify

import (
	"database/sql"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/porrify/porrify_api/database"
	"github.com/porrify/porrify_api/handlers"
)

type Config struct {
	//database
	MysqlUser     string
	MysqlPassword string
	MysqlHost     string
	MysqlPort     string
	MysqlDB       string
}

type properties struct {
	//database
	db *sql.DB
}

var prefixVersion = "/v1"

func initProperties(config *Config) *properties {
	properties := new(properties)
	properties.db = database.OpenDB(config.MysqlUser, config.MysqlPassword,
		config.MysqlHost, config.MysqlPort, config.MysqlDB)
	return properties
}

func Run(config *Config) {
	n := negroni.Classic()
	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", handlers.UserHandler).Methods("GET")
	r.HandleFunc("/users", handlers.AddUserHandler).Methods("POST")
	r.HandleFunc("/users", handlers.UsersHandler).Methods("GET")

	r.HandleFunc("/circuits", handlers.AllCircuitsHandler).Methods("GET")
	r.HandleFunc("/circuits/{id}", handlers.CircuitHandler).Methods("GET")

	r.HandleFunc("/categories/{category}/pilots", handlers.PilotsHandler).Methods("GET")

	r.HandleFunc("/bets", handlers.AddBetHandler).Methods("POST")
	r.HandleFunc("/users/{user_id}/circuits/{circuit_id}/bets", handlers.BetHandler).Methods("GET")

	//TODO better logging, check logrus
	prop := initProperties(config)
	defer prop.db.Close()

	n.UseHandler(corsHandler(r))
	n.Run(":80")
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
