package postegres

import (
	"encoding/json"
	"fmt"
	"github.com/ednailson/miner/internal/domain/entity"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const requestsTableName = "public.requests"
const connectionsTableName = "public.connections"

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
		InsertInto(requestsTableName).
		Cols(
			"data",
			"id",
		).
		Values(
			data,
			req.ID,
		).
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = r.db.Query(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *requestsDataStore) Subscription() (int64, error) {
	rows, err := r.db.Query(fmt.Sprintf("INSERT INTO %s DEFAULT VALUES RETURNING id;", connectionsTableName))
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var id int64
	rows.Next()
	err = rows.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *requestsDataStore) Close() {
	if err := r.db.Close(); err != nil {
		zap.S().Errorf("datastore: failed to close connecation, error: %s", err.Error())
	}
}
