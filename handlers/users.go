package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
		fmt.Println(err.Error())
		return //TODO: Handle erros
	}

	var user models.User

	db, err := sql.Open("mysql", "root:@/porrify")
	if err != nil {
		fmt.Println(err.Error())
		return //TODO: Handle errors
	}

	rows, err := db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Nickname, &user.Avatar)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user": user,
	})
}

// AddUserHandler insert a user in mysql
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	db, err := sql.Open("mysql", "root:@/porrify")

	stmt, err := db.Prepare("INSERT INTO user(email, name, nickname, avatar) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(user.Email, user.Name, user.Nickname, user.Avatar)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
}
