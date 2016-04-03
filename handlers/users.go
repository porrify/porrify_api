package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	// Import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/neomede/porrify_api/models"
)

// UserHandler returns a user
func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User

	db, err := sql.Open("mysql", "root:@/porrify")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	err = db.QueryRow("SELECT * FROM user WHERE id = ?", id).
		Scan(&user.ID, &user.Email, &user.Name, &user.Nickname, &user.Avatar)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(user)
}

// AddUserHandler insert a user in mysql
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		return
	}

	log.Println(user)

	db, err := sql.Open("mysql", "root:@/porrify")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user(id, email, name, nickname, avatar) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
		return
	}
	_, err = stmt.Exec(user.ID, user.Email, user.Name, user.Nickname, user.Avatar)
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
}
