package routes

import (
	"net/http"

	"github.com/1ef7yy/brand_scout/internal/view"
)

func InitHandlers(view view.ViewIFace) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /quotes", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		view.GetQuotes(w, r)
	}))
	mux.Handle("POST /quotes", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		view.CreateQuote(w, r)
	}))
	mux.Handle("GET /quotes/random", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		view.GetRandomQuote(w, r)
	}))
	mux.Handle("DELETE /quotes/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		view.DeleteQuoteByID(w, r)
	}))

	return mux
}
