package record

import (
	"github.com/KonstantinGasser/required"
	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/core/pkg/unit"
)

type Food struct {
	Name     string       `required:"yes"`
	Category FoodCategory `required:"yes"`
	Kcal     unit.Unit    `required:"yes"`
	Carbs    unit.Unit    `required:"yes"`
	Protein  unit.Unit    `required:"yes"`
	Fats     unit.Unit    `required:"yes"`
}

// Valid checks the struct for correctness
//
// if a mandatory field is not in the correct format,
// Valid will indicate this and return false.
func (f Food) Valid() bool {
	if err := required.Atomic(&f); err != nil {
		return false
	}
	return true
}

func FoodFromDTO(food dto.Food) Food {
	return Food{
		Name:     food.Name,
		Category: FoodCategory(food.Category),
		Kcal:     food.Kcal,
		Carbs:    food.Carbs,
		Fats:     food.Fats,
		Protein:  food.Protein,
	}
}

type Recipe struct {
	Name        string
	Ingredients []Food
}
