package usecase

import (
	"github.com/ednailson/miner/internal/domain/entity"
)

type Miner interface {
	Save(req entity.Request) interface{}
}
