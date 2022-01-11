package dao

import "github.com/KonstantinGasser/weeat/core/pkg/unit"

type Food struct {
	ID       int
	Name     string
	Label    string
	Category int
	Kcal     unit.Unit
	Carbs    unit.Unit
	Sugar    unit.Unit
	Protein  unit.Unit
	Fats     unit.Unit
}

type FoodQuery struct {
	ID       int
	Name     string
	Category int
}

type Ingredient struct {
	ID     int
	Amount float64
}

type Recipe struct {
	ID      int
	Name    string
	FoodIDs []Ingredient
}
