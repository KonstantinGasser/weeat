package dao

import "github.com/KonstantinGasser/weeat/core/pkg/unit"

type Food struct {
	Name     string
	Category int
	Kcal     unit.Unit
	Carbs    unit.Unit
	Protein  unit.Unit
	Fats     unit.Unit
}

type Recipe struct {
	Name        string
	Ingredients []Food
}
