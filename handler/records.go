package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/core/services/records"
	"github.com/KonstantinGasser/weeat/pkg/http/json"
	"github.com/KonstantinGasser/weeat/pkg/http/response"
	"github.com/sirupsen/logrus"
)

func HandleInsertFood(recodsvc *records.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqFood dto.Food
		if err := json.FromRequest(r.Body, &reqFood); err != nil {
			logrus.Errorf("[api.InsertFood] could not unmarshal r.Body: %v\n", err)
			response.Err(err, http.StatusBadRequest, "Could not decode data").Write(w)
			return
		}
		defer r.Body.Close()

		inserErr := recodsvc.InsertFood(r.Context(), reqFood)
		if inserErr != nil {
			logrus.Errorf("[api.InsertFood] could not insert food: %v\n", inserErr.Err())
			inserErr.Write(w)
			return
		}
		response.Reply(w).Write(http.StatusCreated, []byte("Food item saved"))
	}
}

func HandleSearchFood(recordsvc *records.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		search := r.URL.Query().Get("q")
		if len(search) == 0 {
			response.Reply(w).Write(http.StatusNotFound, nil)
			return
		}

		items, err := recordsvc.SearchFood(r.Context(), search)
		if err != nil {
			logrus.Errorf("[api.SearchFood] could not lookup food: %v", err.Err())
			err.Write(w)
			return
		}

		response.Reply(w).JSON(http.StatusOK, items)
	}
}

func HandlerInsertRecipe(recordsvc *records.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s works", r.URL)

		b, err := io.ReadAll(r.Body)
		if err != nil {
			logrus.Panic(err)
			return
		}
		defer r.Body.Close()
		fmt.Println(string(b))
	}
}
