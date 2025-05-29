package view

import (
	"encoding/json"
	"net/http"
)

func (v View) CreateQuote(w http.ResponseWriter, r *http.Request) {
}
func (v View) GetQuotes(w http.ResponseWriter, r *http.Request) {
	authorQuery := r.URL.Query().Get("author")

	quotes, err := v.domain.GetQuotes(r.Context(), authorQuery)

	if err != nil {
		v.log.Errorf("error getting quotes: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if quotes == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(quotes)
	if err != nil {
		v.log.Errorf("error marshalling quotes: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		v.log.Errorf("error writing to client: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func (v View) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
}
func (v View) DeleteQuoteByID(w http.ResponseWriter, r *http.Request) {
}
