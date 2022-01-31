package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KonstantinGasser/weeat/core/services/categorysvc"
	"github.com/KonstantinGasser/weeat/pkg/http/response"
	"github.com/sirupsen/logrus"
)

func HandleGetCategories(category *categorysvc.Serivce) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_typeS := r.URL.Query().Get("type")
		if len(_typeS) == 0 {
			response.Err(fmt.Errorf("category type missing"), http.StatusBadRequest, "category type missing").Write(w)
			return
		}

		_type, err := strconv.Atoi(_typeS)
		if err != nil {
			response.Err(fmt.Errorf("category type not a number"), http.StatusBadRequest, "category type not a number")
			return
		}

		categories, caterr := category.Get(r.Context(), _type)
		if caterr != nil {
			logrus.Errorf("[api.GetCategories] could not get categories for type: %d, %v", _type, caterr.Err().Error())
			caterr.Write(w)
			return
		}

		response.Reply(w).JSON(http.StatusOK, map[string]interface{}{
			"data": categories,
		})
	}
}
