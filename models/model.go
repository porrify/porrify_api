package models

// User represents a user
type User struct {
	ID       int    `json:"id"`
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
