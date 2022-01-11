package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/KonstantinGasser/weeat/core/services/recipesvc"
	"github.com/sirupsen/logrus"
)

func HandlerInsertRecipe(recipe *recipesvc.Service) http.HandlerFunc {
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
