package db

import (
	"context"

	"github.com/1ef7yy/brand_scout/internal/models"
)

func (d *DB) CreateQuote(ctx context.Context, req models.CreateQuoteDTO) (models.Quote, error) {
	res, err := d.db.QueryContext(ctx, "INSERT INTO quotes(text, author) VALUES ($1, $2) RETURNING id, text, author", req.Text, req.Author)

	if err != nil {
		d.log.Errorf("error creating quote: %s", err.Error())
		return models.Quote{}, err
	}

	defer res.Close()

	var quote models.Quote

	if res.Next() {
		err = res.Scan(&quote.ID, &quote.Text, &quote.Author)
		if err != nil {
			d.log.Errorf("error scanning into struct: %s", err.Error())
			return models.Quote{}, err
		}
	}
	return quote, nil
}

func (d *DB) GetAllQuotes(ctx context.Context) ([]models.Quote, error) {
	res, err := d.db.QueryContext(ctx, "SELECT id, author, text FROM quotes")

	if err != nil {
		return nil, err
	}
	defer res.Close()

	var quotes []models.Quote

	for res.Next() {
		var quote models.Quote
		err = res.Scan(&quote.ID, &quote.Author, &quote.Text)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

func (d *DB) GetAuthorQuotes(ctx context.Context, author string) ([]models.Quote, error) {
	res, err := d.db.QueryContext(ctx, "SELECT id, author, text FROM quotes WHERE author=$1", author)

	if err != nil {
		return nil, err
	}
	defer res.Close()

	var quotes []models.Quote

	for res.Next() {
		var quote models.Quote
		err = res.Scan(&quote.ID, &quote.Author, &quote.Text)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}


func (d *DB) GetRandomQuote(ctx context.Context) (models.Quote, error) {
	query := `
	SELECT id, text, author FROM quotes ORDER BY RANDOM() LIMIT 1
	`
	res, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return models.Quote{}, err
	}
	defer res.Close()

	var quote models.Quote

	if res.Next() {
		err = res.Scan(&quote.ID, &quote.Text, &quote.Author)
		if err != nil {
			return models.Quote{}, err
		}
	}

	return quote, nil
}


func (d *DB) DeleteQuoteByID(ctx context.Context, id string) (string, error) {
	return "", nil
}
