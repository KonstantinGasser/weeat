package record

import (
	"github.com/KonstantinGasser/required"
	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/core/pkg/category"
	"github.com/KonstantinGasser/weeat/core/pkg/unit"
)

type Food struct {
	Name     string        `required:"yes"`
	Category category.Food `required:"yes"`
	Kcal     unit.Unit     `required:"yes"`
	Carbs    unit.Unit     `required:"yes"`
	Protein  unit.Unit     `required:"yes"`
	Fats     unit.Unit     `required:"yes"`
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
		Category: category.Food(food.Category),
		Kcal:     unit.NewKcal(food.Kcal),
		Carbs:    unit.NewGramm(food.Carbs),
		Fats:     unit.NewGramm(food.Fats),
		Protein:  unit.NewGramm(food.Protein),
	}
}

type Recipe struct {
	Name        string
	Ingredients []Food
}
