package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	// Import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/neomede/porrify_api/models"
)

// BetHandler returns a bet
func BetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	circuitID := vars["circuit_id"]

	db, err := sql.Open("mysql", "root:@/porrify")
	if err != nil {
		log.Println(err.Error())
		return
	}

	rows, err := db.Query("SELECT * FROM bet WHERE user = ? AND circuit = ?", userID, circuitID)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer rows.Close()

	var bets []*models.Bet
	for rows.Next() {
		var bet models.Bet
		err := rows.Scan(&bet.ID, &bet.Category, &bet.Circuit, &bet.Pilot, &bet.Position, &bet.User, &bet.UpdatedAt)
		if err != nil {
			log.Println(err.Error())
			return
		}
		bets = append(bets, &bet)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"bets": bets,
	})
}

// AddBetHandler add or update a bet
func AddBetHandler(w http.ResponseWriter, r *http.Request) {
	var betRace models.BetRace

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&betRace)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db, err := sql.Open("mysql", "root:@/porrify")
	if err != nil {
		log.Println(err.Error())
		return
	}

	stmtInsert, err := db.Prepare("INSERT INTO bet(category, circuit, pilot, position, user, updatedAt) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
		return
	}
	stmtUpdate, err := db.Prepare("UPDATE bet SET pilot = ?, updatedAt = ? WHERE id = ?")
	if err != nil {
		log.Println(err.Error())
		return
	}

	bets := bets(&betRace)
	var currentBet models.Bet
	for _, bet := range bets {
		err = db.QueryRow("SELECT * FROM bet WHERE user = ? AND circuit = ? AND category = ? AND position = ?", bet.User, bet.Circuit, bet.Category, bet.Position).
			Scan(&currentBet.ID, &currentBet.Category, &currentBet.Circuit, &currentBet.Pilot, &currentBet.Position, &currentBet.User, &currentBet.UpdatedAt)
		if err != nil {
			fmt.Println(err.Error())
			_, err = stmtInsert.Exec(bet.Category, bet.Circuit, bet.Pilot, bet.Position, bet.User, time.Now())
			if err != nil {
				log.Println(err.Error())
				return
			}
		} else {
			_, err = stmtUpdate.Exec(bet.Pilot, time.Now(), currentBet.ID)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
}

// bets return a group of bets
func bets(br *models.BetRace) []*models.Bet {
	var bets []*models.Bet
	motogpPoleBet := &models.Bet{
		Category: 1,
		Circuit:  br.Circuit,
		Pilot:    br.MotoGP.Pole,
		Position: 0,
		User:     br.User,
	}
	bets = append(bets, motogpPoleBet)
	motogpFirstBet := &models.Bet{
		Category: 1,
		Circuit:  br.Circuit,
		Pilot:    br.MotoGP.First,
		Position: 1,
		User:     br.User,
	}
	bets = append(bets, motogpFirstBet)
	motogpSecondBet := &models.Bet{
		Category: 1,
		Circuit:  br.Circuit,
		Pilot:    br.MotoGP.Second,
		Position: 2,
		User:     br.User,
	}
	bets = append(bets, motogpSecondBet)
	motogpThirdBet := &models.Bet{
		Category: 1,
		Circuit:  br.Circuit,
		Pilot:    br.MotoGP.Third,
		Position: 3,
		User:     br.User,
	}
	bets = append(bets, motogpThirdBet)
	moto2PoleBet := &models.Bet{
		Category: 2,
		Circuit:  br.Circuit,
		Pilot:    br.Moto2.Pole,
		Position: 0,
		User:     br.User,
	}
	bets = append(bets, moto2PoleBet)
	moto2FirstBet := &models.Bet{
		Category: 2,
		Circuit:  br.Circuit,
		Pilot:    br.Moto2.First,
		Position: 1,
		User:     br.User,
	}
	bets = append(bets, moto2FirstBet)
	moto2SecondBet := &models.Bet{
		Category: 2,
		Circuit:  br.Circuit,
		Pilot:    br.Moto2.Second,
		Position: 2,
		User:     br.User,
	}
	bets = append(bets, moto2SecondBet)
	moto2ThirdBet := &models.Bet{
		Category: 2,
		Circuit:  br.Circuit,
		Pilot:    br.Moto2.Third,
		Position: 3,
		User:     br.User,
	}
	bets = append(bets, moto2ThirdBet)
	moto3PoleBet := &models.Bet{
		Category: 3,
		Circuit:  br.Circuit,
		Pilot:    br.Moto3.Pole,
		Position: 0,
		User:     br.User,
	}
	bets = append(bets, moto3PoleBet)
	moto3FirstBet := &models.Bet{
		Category: 3,
		Circuit:  br.Circuit,
		Pilot:    br.Moto3.First,
		Position: 1,
		User:     br.User,
	}
	bets = append(bets, moto3FirstBet)
	moto3SecondBet := &models.Bet{
		Category: 3,
		Circuit:  br.Circuit,
		Pilot:    br.Moto3.Second,
		Position: 2,
		User:     br.User,
	}
	bets = append(bets, moto3SecondBet)
	moto3ThirdBet := &models.Bet{
		Category: 3,
		Circuit:  br.Circuit,
		Pilot:    br.Moto3.Third,
		Position: 3,
		User:     br.User,
	}
	bets = append(bets, moto3ThirdBet)

	return bets
}
