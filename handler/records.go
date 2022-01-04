package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/KonstantinGasser/lowco/core/services/records"
	"github.com/sirupsen/logrus"
)

func HandleInsertFood(recodsvc *records.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			logrus.Panic(err)
			return
		}
		defer r.Body.Close()
		fmt.Println(string(b))

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
