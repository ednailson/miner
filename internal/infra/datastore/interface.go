package datastore

import "github.com/ednailson/miner/internal/domain/entity"

type DataStore interface {
	Save(req entity.Request) error
	Subscription() (int64, error)
	Close()
}
