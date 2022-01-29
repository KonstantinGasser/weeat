package recipe

import (
	"fmt"
	"strings"

	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/core/pkg/category"
	"github.com/KonstantinGasser/weeat/core/pkg/unit"
)

var (
	ErrDoubleIngredient = fmt.Errorf("recipe's ingredients must be unique")
)

type Ingredient struct {
	ID       int
	Amount   float64
	Name     string
	Category category.Type
	Kcal     unit.Unit
	Carbs    unit.Unit
	Sugar    unit.Unit
	Protein  unit.Unit
	Fats     unit.Unit
}

type Nutritions struct {
	Kcal    unit.Unit
	Carbs   unit.Unit
	Sugar   unit.Unit
	Protein unit.Unit
	Fats    unit.Unit
}

type Recipe struct {
	ID          int
	Name        string
	Category    category.Type
	Label       string
	Nutritions  Nutritions
	Ingredients []Ingredient
}

func New(ID int, name string, category category.Type, ingredients []Ingredient) Recipe {
	return Recipe{
		ID:          ID,
		Name:        name,
		Category:    category,
		Nutritions:  computeNutritions(ingredients),
		Ingredients: ingredients,
	}
}

func computeNutritions(ingredients []Ingredient) Nutritions {
	var (
		kcal    = unit.NewKcal(0)
		carbs   = unit.NewGramm(0)
		sugar   = unit.NewGramm(0)
		protein = unit.NewGramm(0)
		fats    = unit.NewGramm(0)
	)

	for _, item := range ingredients {
		kcal.Add(item.Kcal.Value())
		carbs.Add(item.Carbs.Value())
		sugar.Add(item.Sugar.Value())
		protein.Add(item.Protein.Value())
		fats.Add(item.Fats.Value())
		// item.Kcal.Add(kcal.Value())
		// carbs = item.Carbs.Add(carbs.Value())
		// sugar = item.Sugar.Add(sugar.Value())
		// protein = item.Protein.Add(protein.Value())
		// fats = item.Fats.Add(fats.Value())
	}
	return Nutritions{
		Kcal:    kcal,
		Carbs:   carbs,
		Sugar:   sugar,
		Protein: protein,
		Fats:    fats,
	}
}

func FromDTO(recipe dto.Recipe) (Recipe, error) {

	doubleIngredient := make(map[int]bool)
	for _, ing := range recipe.Ingredients {
		if _, ok := doubleIngredient[ing.ID]; ok {
			return Recipe{}, ErrDoubleIngredient
		}
	}

	var ingredients []Ingredient
	for _, item := range recipe.Ingredients {
		ingredients = append(ingredients, Ingredient{ID: item.ID, Amount: item.Amount})
	}
	return Recipe{
		ID:          recipe.ID,
		Name:        recipe.Name,
		Category:    category.Type(recipe.Category),
		Label:       strings.ToLower(strings.TrimSpace(recipe.Name)),
		Ingredients: ingredients,
	}, nil
}

func (recipe Recipe) ToDTO() dto.Recipe {
	var ing []dto.Ingredient

	for _, item := range recipe.Ingredients {
		ing = append(ing, dto.Ingredient{
			ID:       item.ID,
			Name:     item.Name,
			Amount:   item.Amount,
			Category: int(item.Category),
			Kcal:     int(item.Kcal.Value()),
			Carbs:    item.Carbs.Value(),
			Sugar:    item.Sugar.Value(),
			Protein:  item.Protein.Value(),
			Fats:     item.Fats.Value(),
		})
	}
	return dto.Recipe{
		ID:       recipe.ID,
		Name:     recipe.Name,
		Category: int(recipe.Category),
		Nutritions: dto.Nutritions{
			Kcal:    int(recipe.Nutritions.Kcal.Value()),
			Carbs:   recipe.Nutritions.Carbs.Value(),
			Sugar:   recipe.Nutritions.Sugar.Value(),
			Protein: recipe.Nutritions.Protein.Value(),
			Fats:    recipe.Nutritions.Fats.Value(),
		},
		Ingredients: ing,
	}
}
