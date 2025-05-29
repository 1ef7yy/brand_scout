package domain

import (
	"context"

	"github.com/1ef7yy/brand_scout/internal/models"
)

func (d Domain) CreateQuote(ctx context.Context, author, text string) (quote models.Quote, err error) {
	return models.Quote{}, nil
}
func (d Domain) GetQuotes(ctx context.Context, author string) ([]models.Quote, error) {
	return []models.Quote{}, nil
}
func (d Domain) GetRandomQuote(context.Context) (models.Quote, error) {
	return models.Quote{}, nil
}
func (d Domain) DeleteQuoteByID(context.Context, string) error {
	return nil
}
