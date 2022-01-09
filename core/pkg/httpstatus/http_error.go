package httpstatus

import (
	"net/http"
)

func Created(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
}

func BadRequest(w http.ResponseWriter, errMsg string) {
	http.Error(w, errMsg, http.StatusBadRequest)
}

func InternalError(w http.ResponseWriter, errMsg string) {
	http.Error(w, errMsg, http.StatusInternalServerError)
}
