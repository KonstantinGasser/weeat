package dto

type Food struct {
	Name     string  `json:"name"`
	Category int     `json:"food_cat"`
	Kcal     float64 `json:"kcal"`
	Carbs    float64 `json:"carbs"`
	Sugar    float64 `json:"sugar"`
	Protein  float64 `json:"protein"`
	Fats     float64 `json:"fats"`
}
