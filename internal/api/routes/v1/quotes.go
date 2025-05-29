package v1

import "net/http"

func (v *Router) Quotes() http.Handler {

	mux := http.NewServeMux()

	mux.Handle("POST /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v.View.CreateQuote(w, r)
	}))

	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v.View.GetQuotes(w, r)
	}))

	mux.Handle("GET /random", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v.View.GetRandomQuote(w, r)
	}))

	mux.Handle("DELETE /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v.View.DeleteQuoteByID(w, r)
	}))

	return http.StripPrefix("/quotes", mux)
}
