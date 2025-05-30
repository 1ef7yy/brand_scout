package view

import (
	"net/http"

	"github.com/1ef7yy/brand_scout/internal/domain"
	"github.com/1ef7yy/brand_scout/pkg/logger"
)

type View struct {
	log    logger.Logger
	domain domain.DomainIFace
}

type ViewIFace interface {
	CreateQuote(w http.ResponseWriter, r *http.Request)
	GetQuotes(w http.ResponseWriter, r *http.Request)
	GetRandomQuote(w http.ResponseWriter, r *http.Request)
	DeleteQuoteByID(w http.ResponseWriter, r *http.Request)
}

func NewView(log logger.Logger) (ViewIFace, error) {
	domain, err := domain.NewDomain(log)
	if err != nil {
		log.Errorf("error creating domain: %s", err.Error())
		return nil, err
	}
	return &View{
		log:    log,
		domain: domain,
	}, nil
}
