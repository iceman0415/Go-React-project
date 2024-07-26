package db

// Person contains information about a person.
type Person struct {
	Name    string  `json:"name,omitempty"`
	Age     int     `json:"age,omitempty"`
	Balance float32 `json:"balance,omitempty"`
	Email   string  `json:"email"`
	Address string  `json:"address,omitempty"`
}
