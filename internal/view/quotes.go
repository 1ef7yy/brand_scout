package view

import (
	"encoding/json"
	"net/http"

	customErrors "github.com/1ef7yy/brand_scout/internal/errors"
	"github.com/1ef7yy/brand_scout/internal/models"
)

func (v *View) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var createReq models.CreateQuoteDTO

	err := json.NewDecoder(r.Body).Decode(&createReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		v.log.Errorf("error scanning JSON into struct: %s", err.Error())
		return
	}

	quote, err := v.domain.CreateQuote(r.Context(), createReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		v.log.Errorf("error creating quote: %s", err.Error())
		return
	}

	resp, err := json.Marshal(quote)
	if err != nil {
		v.log.Errorf("error marshalling struct: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(resp)
	if err != nil {
		v.log.Errorf("error writing to client: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func (v *View) GetQuotes(w http.ResponseWriter, r *http.Request) {
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
func (v *View) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := v.domain.GetRandomQuote(r.Context())
	if err != nil {
		v.log.Errorf("error getting random quote: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if quote.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(quote)
	if err != nil {
		v.log.Errorf("error marshalling struct: %s", err.Error())
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
func (v *View) DeleteQuoteByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	deletedID, err := v.domain.DeleteQuoteByID(r.Context(), id)
	if err == customErrors.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		v.log.Errorf("error deleting quote: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if deletedID == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(deletedID)

	if err != nil {
		v.log.Errorf("error deleting quote by id %s: %s", id, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		v.log.Errorf("error writing to client: %s", err.Error())
		return
	}
}
