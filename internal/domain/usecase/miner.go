package usecase

import (
	"github.com/ednailson/miner/internal/domain/entity"
	"github.com/ednailson/miner/internal/infra/datastore"
	"go.uber.org/zap"
)

const minerVersion = "cgminer/4.10.0"

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
	case entity.SubscribeMethod:
		switch params := req.Params.(type) {
		case []interface{}:
			if len(params) != 1 || params[0] != minerVersion {
				return entity.NewFail(&req.ID, entity.ErrorInvalidParams())
			}
		case map[string]interface{}:
			if params["version"] != minerVersion {
				return entity.NewFail(&req.ID, entity.ErrorInvalidParams())
			}
		default:
			return entity.NewFail(&req.ID, entity.ErrorInvalidParams())
		}
		id, err := u.ds.Subscription()
		if err != nil {
			zap.S().Errorf("usecase: subscription error: %s", err.Error())
			return entity.NewFail(&req.ID, entity.ErrorServer())
		}
		result = entity.NewSubscription(id)
	default:
		return entity.NewFail(&req.ID, entity.ErrorMethodNotFound())
	}

	err := u.ds.Save(req)
	if err != nil {
		zap.S().Errorf("usecase: failed to save on datastore, error: %s", err.Error())
		return entity.NewFail(&req.ID, entity.ErrorServer())
	}

	return entity.NewSuccess(req.ID, result)
}
