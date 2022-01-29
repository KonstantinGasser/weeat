package recipesvc

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/weeat/core/dao"
	"github.com/KonstantinGasser/weeat/core/domain/recipe"
	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/core/pkg/category"
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
		ID:       recipeItem.ID,
		Name:     recipeItem.Name,
		Label:    recipeItem.Label,
		Category: int(recipeItem.Category),
	}
	for _, item := range recipeItem.Ingredients {
		daoItem.FoodIDs = append(daoItem.FoodIDs, dao.Ingredient{ID: item.ID, Amount: item.Amount})
	}
	if err := svc.repo.InsertRecipe(ctx, daoItem); err != nil {
		return response.Err(err, http.StatusInternalServerError, "could not store recipe")
	}

	return nil
}

func (svc Service) Search(ctx context.Context, query string, limit int) ([]dto.Recipe, response.RespErr) {

	foundRecipes, err := svc.repo.SearchRecipe(ctx, query, limit)
	if err != nil {
		return nil, response.Err(err, http.StatusBadRequest, "Could not search for recipes")
	}

	var recipes []dto.Recipe
	for _, found := range foundRecipes {

		var ing []recipe.Ingredient
		for _, i := range found.FoodIDs {
			ing = append(ing, recipe.Ingredient{
				ID:       i.ID,
				Amount:   i.Amount,
				Name:     i.Food.Name,
				Category: category.Type(i.Food.Category),
				Kcal:     i.Food.Kcal,
				Carbs:    i.Food.Carbs,
				Sugar:    i.Food.Sugar,
				Protein:  i.Food.Protein,
				Fats:     i.Food.Fats,
			})
		}

		rec := recipe.New(found.ID, found.Name, category.Type(found.Category), ing).ToDTO()
		recipes = append(recipes, rec)
	}
	return recipes, nil
}
