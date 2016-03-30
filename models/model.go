package models

// User represents a user
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// Category represents a MotoGP category
type Category int

// Pilot represents a pilot
type Pilot struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Number   int      `json:"number"`
	Category Category `json:"category"`
}

// Circuit represents a circuit
type Circuit struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Day     string `json:"day"` //TODO: Get day properly
}

// BetRace represents a bet for the race
type BetRace struct {
	MotoGP  BetCategory `json:"motogp"`
	Moto2   BetCategory `json:"moto2"`
	Moto3   BetCategory `json:"moto3"`
	Circuit int         `json:"circuit"`
	User    string      `json:"user"`
}

// BetCategory represents a bet for a category
type BetCategory struct {
	Pole   int `json:"pole"`
	First  int `json:"first"`
	Second int `json:"second"`
	Third  int `json:"third"`
}

// Bet represents a single bet
type Bet struct {
	ID        int      `json:"id"`
	Category  Category `json:"category"`
	Circuit   int      `json:"circuit"`
	Pilot     int      `json:"pilot"`
	Position  int      `json:"position"`
	User      string   `json:"user"`
	UpdatedAt string   `json:"updated_at"` //TODO: Get date properly
}
