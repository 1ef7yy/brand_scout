package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/1ef7yy/brand_scout/pkg/logger"
	_ "github.com/lib/pq" // for postgres driver
)

type DB struct {
	log logger.Logger
	db  *sql.DB
}

func New(log logger.Logger) (*DB, error) {
	postgresConn, ok := os.LookupEnv("POSTGRES_CONN")

	if !ok {
		log.Errorf("could not resolve POSTGRES_CONN")
		return nil, fmt.Errorf("could not resolve POSTGRES_CONN")
	}

	db, err := sql.Open("postgres", postgresConn)

	if err != nil {
		log.Errorf("error opening db: %s", err.Error())
		return nil, err
	}

	log.Info("db connection opened")

	return &DB{
		log: log,
		db:  db,
	}, nil
}

func (d *DB) Init() error {
	_, err := d.db.Exec(`
		CREATE TABLE IF NOT EXISTS quotes (
			id BIGSERIAL PRIMARY KEY,
			text TEXT NOT NULL,
			author TEXT NOT NULL
		)`,
	)
	d.log.Info("initialized database")
	return err
}
