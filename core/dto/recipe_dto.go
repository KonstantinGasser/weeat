package dto

type Ingredient struct {
	ID       int     `json:"id"`
	Amount   float64 `json:"amount"`
	Name     string  `json:"name"`
	Category int     `json:"category"`
	Kcal     int     `json:"kcal"`
	Carbs    float64 `json:"carbs"`
	Sugar    float64 `json:"sugar"`
	Protein  float64 `json:"protein"`
	Fats     float64 `json:"fats"`
}

type Nutritions struct {
	Kcal    int     `json:"kcal"`
	Carbs   float64 `json:"carbs"`
	Sugar   float64 `json:"sugar"`
	Protein float64 `json:"protein"`
	Fats    float64 `json:"fats"`
}

type Recipe struct {
	ID          int          `json:"id,omitempty"`
	Name        string       `json:"name"`
	Category    int          `json:"category"`
	Label       string       `json:"label,omitempty"`
	Nutritions  Nutritions   `json:"nutritions"`
	Ingredients []Ingredient `json:"ingredients"`
}
