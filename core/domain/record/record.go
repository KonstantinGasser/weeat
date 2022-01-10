package record

import (
	"fmt"
	"strings"

	"github.com/KonstantinGasser/required"
	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/core/pkg/category"
	"github.com/KonstantinGasser/weeat/core/pkg/unit"
	"github.com/pkg/errors"
)

var (
	ErrMissingData      = fmt.Errorf("data for record.Food is not compliant with struct requirements")
	ErrSugarBiggerCarbs = fmt.Errorf("record.Food.Sugar is bigger than total carbs; this is a contradiction")
)

type Food struct {
	Name string `required:"yes"`
	// Label is the lower-case representation for the Food
	// it is used for fast access search of items
	Label    string
	Category category.Food `required:"yes"`
	Kcal     unit.Unit     `required:"yes"`
	Carbs    unit.Unit     `required:"yes"`
	Sugar    unit.Unit     `required:"yes"`
	Protein  unit.Unit     `required:"yes"`
	Fats     unit.Unit     `required:"yes"`
}

// Valid checks the struct for correctness
//
// if a mandatory field is not in the correct format,
// Valid will indicate this and return an error.
func (f Food) Valid() error {
	if err := required.Atomic(&f); err != nil {
		return errors.Wrap(err, ErrMissingData.Error())
	}

	// surge is a subset of carbs and thereby cannot be bigger
	// than the provided carbs value
	if f.Sugar.Value() > f.Carbs.Value() {
		return ErrSugarBiggerCarbs
	}
	return nil
}

func FoodFromDTO(food dto.Food) Food {
	return Food{
		Name:     food.Name,
		Label:    strings.ToLower(food.Name),
		Category: category.Food(food.Category),
		Kcal:     unit.NewKcal(food.Kcal),
		Carbs:    unit.NewGramm(food.Carbs),
		Sugar:    unit.NewGramm(food.Sugar),
		Fats:     unit.NewGramm(food.Fats),
		Protein:  unit.NewGramm(food.Protein),
	}
}

type Recipe struct {
	Name        string
	Ingredients []Food
}
