package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/core/pkg/httpstatus"
	"github.com/KonstantinGasser/weeat/core/pkg/json"
	"github.com/KonstantinGasser/weeat/core/services/records"
	"github.com/sirupsen/logrus"
)

func HandleInsertFood(recodsvc *records.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqFood dto.Food
		if err := json.FromRequest(r.Body, &reqFood); err != nil {
			logrus.Errorf("[api.InsertFood] could not unmarshal r.Body: %v\n", err)
			httpstatus.BadRequest(w, "malformed JSON body")
			return
		}
		defer r.Body.Close()

		fmt.Printf("%+v", reqFood)

		inserErr := recodsvc.InsertFood(r.Context(), reqFood)
		if inserErr != nil {
			logrus.Errorf("[api.InsertFood] could not insert food: %v\n", inserErr)
			httpstatus.InternalError(w, inserErr.Error())
			return
		}
		httpstatus.Created(w)
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
