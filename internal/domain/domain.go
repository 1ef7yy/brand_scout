package domain

import (
	"context"

	"github.com/1ef7yy/brand_scout/internal/models"
	"github.com/1ef7yy/brand_scout/pkg/logger"
)

type Domain struct {
	log logger.Logger
}

type DomainIFace interface {
	CreateQuote(ctx context.Context, author, text string) (quote models.Quote, err error)
	GetQuotes(ctx context.Context, author string) ([]models.Quote, error)
	GetRandomQuote(context.Context) (models.Quote, error)
	DeleteQuoteByID(context.Context, string) error
}

func NewDomain(log logger.Logger) (DomainIFace, error) {
	return Domain{
		log: log,
	}, nil
}
