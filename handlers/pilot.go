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
	"github.com/porrify/porrify_api/models"
)

// PilotsHandler returns all pilots
func PilotsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category, err := strconv.Atoi(vars["category"])
	if err != nil {
		log.Println(err.Error())
		return
	}

	var pilots []*models.Pilot

	db, err := sql.Open("mysql", "root:@/porrify")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM pilot WHERE category = ?", category)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var pilot models.Pilot
		err := rows.Scan(&pilot.ID, &pilot.Name, &pilot.Number, &pilot.Category)
		if err != nil {
			log.Println(err.Error())
			return
		}

		pilots = append(pilots, &pilot)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"pilots": pilots,
	})
}
