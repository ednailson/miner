package postegres

import (
	"encoding/json"
	"github.com/ednailson/miner/internal/domain/entity"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const tableName = "requests"

type requestsDataStore struct {
	db *sqlx.DB
}

func NewRequestDataStore(db *sqlx.DB) *requestsDataStore {
	return &requestsDataStore{db: db}
}

func (r *requestsDataStore) Save(req entity.Request) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	query, args := sqlbuilder.
		InsertInto(tableName).
		Cols(
			"data",
			"id",
		).
		Values(
			data,
			req.ID,
		).
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *requestsDataStore) Close() {
	if err := r.db.Close(); err != nil {
		zap.S().Errorf("datastore: failed to close connecation, error: %s", err.Error())
	}
}
