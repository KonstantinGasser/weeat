package recipesvc

import (
	"context"

	"github.com/KonstantinGasser/weeat/core/dao"
)

type RecipeRepo interface {
	InsertRecipe(ctx context.Context, recipe dao.Recipe) (int, error)
}
