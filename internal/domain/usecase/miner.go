package usecase

import (
	"github.com/ednailson/miner/internal/domain/entity"
	"github.com/ednailson/miner/internal/infra/datastore"
)

type useCaseMiner struct {
	ds datastore.DataStore
}

func NewMiner(ds datastore.DataStore) *useCaseMiner {
	return &useCaseMiner{ds: ds}
}

func (u *useCaseMiner) Save(req entity.Request) interface{} {
	var result interface{}
	switch req.Method {
	case entity.AuthorizeMethod:
		result = true
	case entity.SubscriptionMethod:
		result = entity.Subscription{}
	default:
		return entity.NewFail(&req.ID, entity.ErrorMethodNotFound())
	}

	err := u.ds.Save(req)
	if err != nil {
		return entity.NewFail(&req.ID, entity.ErrorServer())
	}

	return entity.NewSuccess(req.ID, result)
}
