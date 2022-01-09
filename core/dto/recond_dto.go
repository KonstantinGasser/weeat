package dto

type Food struct {
	Name     string  `json:"name"`
	Category int     `json:"food_cat"`
	Kcal     float64 `json:"kcal"`
	Carbs    float64 `json:"carbs"`
	Protein  float64 `json:"protein"`
	Fats     float64 `json:"fats"`
}
