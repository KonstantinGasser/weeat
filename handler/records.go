package handler

import (
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/lowco/core/services/records"
)

func HandleInsertFood(recodsvc *records.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "InsertFood works")
	}
}
