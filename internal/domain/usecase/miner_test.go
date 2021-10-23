package usecase

import (
	"fmt"
	"github.com/ednailson/miner/internal/domain/entity"
	"github.com/ednailson/miner/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestMiner(t *testing.T) {
	zap.NewNop()
	t.Run("authorize successfully", func(t *testing.T) {
		ds := new(mocks.DataStore)
		request := entity.FakeRequest(entity.AuthorizeMethod, nil)
		ds.On("Save", request).Return(nil)
		sut := NewMiner(ds)

		response := sut.Save(request)

		assert.Equal(t, response, entity.NewSuccess(request.ID, true))
	})
	t.Run("subscribe successfully", func(t *testing.T) {
		ds := new(mocks.DataStore)
		request := entity.FakeRequest(entity.SubscribeMethod, []string{minerVersion})
		id := int64(6489)
		ds.On("Subscription").Return(id, nil)
		ds.On("Save", request).Return(nil)
		sut := NewMiner(ds)

		response := sut.Save(request)

		assert.Equal(t, response, entity.NewSuccess(request.ID, entity.NewSubscription(id)))
	})
	t.Run("invalid method", func(t *testing.T) {
		ds := new(mocks.DataStore)
		request := entity.FakeRequest("invalid.method", nil)
		sut := NewMiner(ds)

		response := sut.Save(request)

		assert.Equal(t, response, entity.NewFail(&request.ID, entity.ErrorMethodNotFound()))
	})
	t.Run("error on request method in datastore", func(t *testing.T) {
		ds := new(mocks.DataStore)
		request := entity.FakeRequest(entity.AuthorizeMethod, nil)
		ds.On("Save", request).Return(fmt.Errorf(""))
		sut := NewMiner(ds)

		response := sut.Save(request)

		assert.Equal(t, response, entity.NewFail(&request.ID, entity.ErrorServer()))
	})
	t.Run("error on subscription method in datastore", func(t *testing.T) {
		ds := new(mocks.DataStore)
		request := entity.FakeRequest(entity.SubscribeMethod, []string{minerVersion})
		ds.On("Subscription").Return(int64(0), fmt.Errorf(""))
		sut := NewMiner(ds)

		response := sut.Save(request)

		assert.Equal(t, response, entity.NewFail(&request.ID, entity.ErrorServer()))
	})
}
