package food

import (
	"fmt"
	"strings"

	"github.com/KonstantinGasser/required"
	"github.com/KonstantinGasser/weeat/core/dao"
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
	ID   int
	Name string `required:"yes"`
	// Label is the lower-case representation for the Food Name
	// it is used for fast access search of items
	Label    string
	Category category.Type `required:"yes"`
	Kcal     unit.Unit     `required:"yes"`
	Carbs    unit.Unit     `required:"yes"`
	Sugar    unit.Unit     `required:"yes"`
	Protein  unit.Unit     `required:"yes"`
	Fats     unit.Unit     `required:"yes"`
}

func (f Food) Scale(scaler int) Food {
	return Food{
		ID:       f.ID,
		Name:     f.Name,
		Category: f.Category,
		Kcal:     f.Kcal.Scale(scaler),
		Carbs:    f.Carbs.Scale(scaler),
		Sugar:    f.Sugar.Scale(scaler),
		Protein:  f.Protein.Scale(scaler),
		Fats:     f.Fats.Scale(scaler),
	}
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
		Category: category.Type(food.Category),
		Kcal:     unit.NewKcal(food.Kcal),
		Carbs:    unit.NewGramm(food.Carbs),
		Sugar:    unit.NewGramm(food.Sugar),
		Fats:     unit.NewGramm(food.Fats),
		Protein:  unit.NewGramm(food.Protein),
	}
}

func FoodFromDAO(food dao.Food) Food {
	return Food{
		ID:       food.ID,
		Name:     food.Name,
		Label:    food.Name,
		Category: category.Type(food.Category),
		Kcal:     food.Kcal,
		Carbs:    food.Carbs,
		Sugar:    food.Sugar,
		Fats:     food.Fats,
		Protein:  food.Protein,
	}
}
