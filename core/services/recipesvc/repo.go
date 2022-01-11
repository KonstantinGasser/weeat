package recipesvc

import (
	"context"

	"github.com/KonstantinGasser/weeat/core/dao"
)

type RecipeRepo interface {
	InsertRecipe(ctx context.Context, recipe dao.Recipe) error
	MapFoodToRecipe(ctx context.Context, recipe int, foods ...dao.Ingredient) error
}
