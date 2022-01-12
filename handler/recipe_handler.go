package handler

import (
	"net/http"

	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/core/services/recipesvc"
	"github.com/KonstantinGasser/weeat/pkg/http/json"
	"github.com/KonstantinGasser/weeat/pkg/http/response"
	"github.com/sirupsen/logrus"
)

func HandlerInsertRecipe(recipe *recipesvc.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqBody dto.Recipe
		if err := json.FromRequest(r.Body, &reqBody); err != nil {
			logrus.Errorf("[api.InsertRecipe] could not unmarshal r.Body: %v\n", err)
			response.Err(err, http.StatusBadRequest, "Could not unmarshal data").Write(w)
			return
		}

		if err := recipe.InsertRecipe(r.Context(), reqBody); err != nil {
			logrus.Errorf("[api.InsertRecipe] could not store recipe: %v\n", err.Err())
			err.Write(w)
			return
		}
		response.Reply(w).Write(http.StatusCreated, []byte("Save your recipe"))
	}
}
