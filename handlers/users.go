package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	// Import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/neomede/porrify_api/models"
)

// UserHandler returns a user
func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		return
	}

	var user models.User

	db, err := sql.Open("mysql", "root:@/porrify")
	if err != nil {
		log.Println(err.Error())
		return
	}

	rows, err := db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Nickname, &user.Avatar)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(user)
}

// AddUserHandler insert a user in mysql
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println(err.Error())
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Println(err.Error())
		return
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Println(err.Error())
			return
		}
	}

	db, err := sql.Open("mysql", "root:@/porrify")

	stmt, err := db.Prepare("INSERT INTO user(email, name, nickname, avatar) VALUES(?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
		return
	}
	_, err = stmt.Exec(user.Email, user.Name, user.Nickname, user.Avatar)
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
}
