package handler

import (
	"net/http"

	"github.com/KonstantinGasser/weeat/core/services/verify"
)

func HandleOpenSSE(verifyer verify.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("connection", "keep-alive")
		w.Header().Set("Content-T", "stream/text")

		for 
	}
}
