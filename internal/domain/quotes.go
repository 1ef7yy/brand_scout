package domain

import (
	"context"

	"github.com/1ef7yy/brand_scout/internal/models"
)

func (d *Domain) CreateQuote(ctx context.Context, req models.CreateQuoteDTO) (models.Quote, error) {
	quote, err := d.db.CreateQuote(ctx, req)
	if err != nil {
		d.log.Errorf("error creating quote: %s", err.Error())
		return models.Quote{}, err
	}

	return quote, nil
}
func (d *Domain) GetQuotes(ctx context.Context, author string) ([]models.Quote, error) {
	if author == "" {
		quotes, err := d.db.GetAllQuotes(ctx)

		if err != nil {
			d.log.Errorf("error getting all quotes: %s", err.Error())
			return nil, err
		}
		return quotes, nil
	} else {
		quotes, err := d.db.GetAuthorQuotes(ctx, author)
		if err != nil {
			d.log.Errorf("error getting author %s quotes: %s", author, err.Error())
			return nil, err
		}
		return quotes, nil
	}
}
func (d *Domain) GetRandomQuote(ctx context.Context) (models.Quote, error) {
	quote, err := d.db.GetRandomQuote(ctx)


	if err != nil {
		d.log.Errorf("error getting random quote: %s", err.Error())
	}
	return quote, err
}
func (d *Domain) DeleteQuoteByID(ctx context.Context, id string) (string, error) {
	deletedID, err := d.db.DeleteQuoteByID(ctx, id)
	if err != nil {
		d.log.Errorf("error deleting quote by id %s: %s", id, err.Error())
		return "", err
	}

	return deletedID, nil
}
