package record

import "github.com/KonstantinGasser/weeat/core/pkg/unit"

type Food struct {
	Name     string
	Category FoodCategory
	Kcal     unit.Unit
	Carbs    unit.Unit
	Protein  unit.Unit
	Fats     unit.Unit
}

type Recipe struct {
	Name        string
	Ingredients []Food
}
