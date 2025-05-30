package domain

import (
	"context"

	"github.com/1ef7yy/brand_scout/internal/models"
	db "github.com/1ef7yy/brand_scout/internal/storage"
	"github.com/1ef7yy/brand_scout/pkg/logger"
)

type Domain struct {
	log logger.Logger
	db  *db.DB
}

type DomainIFace interface {
	CreateQuote(ctx context.Context, createReq models.CreateQuoteDTO) (quote models.Quote, err error)
	GetQuotes(ctx context.Context, author string) ([]models.Quote, error)
	GetRandomQuote(context.Context) (models.Quote, error)
	DeleteQuoteByID(context.Context, string) (string, error)
}

func NewDomain(log logger.Logger) (DomainIFace, error) {
	db, err := db.New(log)
	if err != nil {
		log.Errorf("error creating db instance: %s", err.Error())
		return nil, err
	}

	err = db.Init()
	if err != nil {
		log.Errorf("error initializing db: %s", err.Error())
		return nil, err
	}
	return &Domain{
		log: log,
		db:  db,
	}, nil
}
