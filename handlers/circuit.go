package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	// Import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/neomede/porrify_api/models"
)

// CircuitHandler returns a circuit
func CircuitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		return
	}

	var circuit models.Circuit

	db, err := sql.Open("mysql", "root:@/porrify")
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = db.QueryRow("SELECT * FROM circuit WHERE id = ?", id).
		Scan(&circuit.ID, &circuit.Name, &circuit.Country, &circuit.Day)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(circuit)
}

// AllCircuitsHandler returns all circuits
func AllCircuitsHandler(w http.ResponseWriter, r *http.Request) {
	var circuits []*models.Circuit

	db, err := sql.Open("mysql", "root:@/porrify")
	if err != nil {
		log.Println(err.Error())
		return
	}

	rows, err := db.Query("SELECT * FROM circuit")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var circuit models.Circuit
		err := rows.Scan(&circuit.ID, &circuit.Name, &circuit.Country, &circuit.Day)
		if err != nil {
			log.Println(err.Error())
			return
		}
		circuits = append(circuits, &circuit)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"circuits": circuits,
	})
}
