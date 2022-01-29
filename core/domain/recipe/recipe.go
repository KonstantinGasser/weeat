package recipe

import (
	"fmt"
	"strings"

	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/core/pkg/category"
)

var (
	ErrDoubleIngredient = fmt.Errorf("recipe's ingredients must be unique")
)

type Ingredient struct {
	ID     int
	Amount float64
}

type Recipe struct {
	ID          int
	Name        string
	Category    category.Type
	Label       string
	Ingredients []Ingredient
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
