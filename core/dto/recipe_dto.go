package dto

type Ingredient struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
}

type Recipe struct {
	ID          int          `json:"id,omitempty"`
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}
