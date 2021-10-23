package database

import (
	"database/sql"
	"github.com/ednailson/miner/internal/infra/datastore"
	"github.com/ednailson/miner/internal/infra/datastore/postegres"
	"github.com/jmoiron/sqlx"
	"os"
)

func Setup() (datastore.DataStore, error) {
	const driverName = "postgres"
	db, err := sql.Open(driverName, os.Getenv("DATABASE_HOST"))
	if err != nil {
		return nil, err
	}

	dbx := sqlx.NewDb(db, driverName)

	if err = dbx.Ping(); err != nil {
		return nil, err
	}

	return postegres.NewRequestDataStore(dbx), nil
}
