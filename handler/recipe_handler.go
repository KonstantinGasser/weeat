package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/core/services/recipesvc"
	"github.com/KonstantinGasser/weeat/pkg/http/json"
	"github.com/KonstantinGasser/weeat/pkg/http/response"
	"github.com/sirupsen/logrus"
)

func HandleInsertRecipe(recipe *recipesvc.Service) http.HandlerFunc {
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

func HandleSearchRecipe(recipe *recipesvc.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		search := strings.ToLower(r.URL.Query().Get("q"))
		if len(search) == 0 {
			response.Reply(w).JSON(http.StatusOK, nil)
			return
		}

		limit, limiterr := strconv.Atoi(r.URL.Query().Get("l"))
		if limiterr != nil {
			logrus.Errorf("[api.SearchRecipe] requested limit not a number: got %s\n", limiterr)
			response.Err(limiterr, http.StatusBadRequest, "Search limit not a number").Write(w)
			return
		}

		items, err := recipe.Search(r.Context(), search, limit)
		if err != nil {
			logrus.Errorf("[api.SearchRecipe] could not lookup food: %v", err.Err())
			err.Write(w)
			return
		}

		response.Reply(w).JSON(http.StatusOK, items)
	}
}
