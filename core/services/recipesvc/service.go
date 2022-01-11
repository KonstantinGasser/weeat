package recipesvc

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/weeat/core/dao"
	"github.com/KonstantinGasser/weeat/core/domain/recipe"
	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/pkg/http/response"
)

type Service struct {
	repo RecipeRepo
}

func New(repo RecipeRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc Service) InsertRecipe(ctx context.Context, r dto.Recipe) response.RespErr {

	recipeItem, err := recipe.FromDTO(r)
	if err != nil {
		return response.Err(err, http.StatusBadRequest, "Ingredients must be unique for a recipe")
	}

	daoItem := dao.Recipe{
		ID:   recipeItem.ID,
		Name: recipeItem.Name,
	}
	for _, item := range recipeItem.Ingredients {
		daoItem.FoodIDs = append(daoItem.FoodIDs, dao.Ingredient{ID: item.ID, Amount: item.Amount})
	}
	if err := svc.repo.InsertRecipe(ctx, daoItem); err != nil {
		return response.Err(err, http.StatusInternalServerError, "could not store recipe")
	}

	// if err := svc.repo.MapFoodToRecipe(ctx, newID, daoItem.FoodIDs...); err != nil {
	// 	return response.Err(err, http.StatusInternalServerError, "could not store recipe")
	// }

	return nil
}
