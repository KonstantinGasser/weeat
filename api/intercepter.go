package api

import "net/http"

type methodIntercepter func(http.HandlerFunc) http.HandlerFunc

func MethodGet(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// server next func (can be another intercepter or the handler func)
		next(w, r)
	}
}

func MethodPut(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// server next func (can be another intercepter or the handler func)
		next(w, r)
	}
}

func MethodPost(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// server next func (can be another intercepter or the handler func)
		next(w, r)
	}
}

func MethodDelete(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// server next func (can be another intercepter or the handler func)
		next(w, r)
	}
}
