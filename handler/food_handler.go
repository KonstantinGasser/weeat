package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/KonstantinGasser/weeat/core/domain/verification"
	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/core/services/foodsvc"
	"github.com/KonstantinGasser/weeat/core/services/verify"
	"github.com/KonstantinGasser/weeat/pkg/http/json"
	"github.com/KonstantinGasser/weeat/pkg/http/response"
	"github.com/sirupsen/logrus"
)

func HandleInsertFood(food *foodsvc.Service, verifyer *verify.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqFood dto.Food
		if err := json.FromRequest(r.Body, &reqFood); err != nil {
			logrus.Errorf("[api.InsertFood] could not unmarshal r.Body: %v\n", err)
			response.Err(err, http.StatusBadRequest, "Could not decode data").Write(w)
			return
		}
		defer r.Body.Close()

		inserErr := food.Insert(r.Context(), reqFood)
		if inserErr != nil {
			logrus.Errorf("[api.InsertFood] could not insert food: %v\n", inserErr.Err())
			inserErr.Write(w)
			return
		}
		// reqFood.ID is wrong !! will be empty since the id is assigned by the database
		verifyer.Trigger(verification.FoodEvent{ID: reqFood.ID})
		response.Reply(w).Write(http.StatusCreated, []byte("Food item saved"))
	}
}

func HandleSearchFood(food *foodsvc.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		search := strings.ToLower(r.URL.Query().Get("q"))
		if len(search) == 0 {
			response.Reply(w).JSON(http.StatusOK, nil)
			return
		}

		limit, limiterr := strconv.Atoi(r.URL.Query().Get("l"))
		if limiterr != nil {
			logrus.Errorf("[api.SearchFood] requested limit not a number: got %s\n", limiterr)
			response.Err(limiterr, http.StatusBadRequest, "Search limit not a number").Write(w)
			return
		}

		items, err := food.Search(r.Context(), search, limit)
		if err != nil {
			logrus.Errorf("[api.SearchFood] could not lookup food: %v", err.Err())
			err.Write(w)
			return
		}

		response.Reply(w).JSON(http.StatusOK, items)
	}
}

func HandleGetFood(food *foodsvc.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		itemID := r.URL.Query().Get("id")
		if len(itemID) == 0 {
			response.Err(fmt.Errorf("item id missing"), http.StatusBadRequest, "item id missing").Write(w)
			return
		}
		s := r.URL.Query().Get("scaler")
		if len(s) == 0 {
			response.Err(fmt.Errorf("amount in query missing"), http.StatusBadRequest, "Amount cannot be zero").Write(w)
			return
		}

		scaler, convErr := strconv.Atoi(s)
		if convErr != nil {
			response.Err(convErr, http.StatusBadRequest, "Amount must be a valid number").Write(w)
			return
		}

		items, err := food.Get(r.Context(), itemID, scaler)
		if err != nil {
			logrus.Errorf("[api.SearchFood] could not lookup food: %v", err.Err())
			err.Write(w)
			return
		}

		response.Reply(w).JSON(http.StatusOK, items)
	}
}
