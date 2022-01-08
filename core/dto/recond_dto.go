package dto

import "github.com/KonstantinGasser/weeat/core/pkg/unit"

type Food struct {
	Name     string    `json:"name"`
	Category int       `json:"food_cat`
	Kcal     unit.Unit `json:"kcal"`
	Carbs    unit.Unit `json:"carbs"`
	Protein  unit.Unit `json:"protein"`
	Fats     unit.Unit `json:"fats"`
}
