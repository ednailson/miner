package server

import (
	"bufio"
	"fmt"
	"github.com/ednailson/miner/internal/domain/entity"
	"github.com/ednailson/miner/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"net"
	"testing"
	"time"
)

var breakLine = []byte("\n")

func TestServer(t *testing.T) {
	zap.NewNop()
	addr := ":9999"
	t.Run("authorize request successfully", func(t *testing.T) {
		uc := new(mocks.Miner)
		id := uint64(543)
		uc.On("Save", mock.Anything).Return(entity.NewSuccess(id, true))

		sut := NewServer(addr, uc)
		go sut.Start()
		time.Sleep(10 * time.Millisecond)
		defer sut.Close()

		client, err := net.Dial("tcp", addr)
		assert.NoError(t, err)
		defer client.Close()

		go func() {
			_, err = client.Write(fakeAuthorize(id))
			assert.NoError(t, err)
			_, err = client.Write(breakLine)
			assert.NoError(t, err)
		}()

		assertResponse(t, fakeAuthorizeResponse(id), client)

	})
}

func assertResponse(t *testing.T, match []byte, conn net.Conn) {
	reader := bufio.NewReader(conn)
	l, _, err := reader.ReadLine()
	assert.NoError(t, err)
	assert.Equal(t, l, match)

}

func fakeAuthorizeResponse(id uint64) []byte {
	return []byte(fmt.Sprintf(`{"jsonrpc":"2.0","result":true,"id":%d}`, id))
}

func fakeAuthorize(id uint64) []byte {
	return []byte(fmt.Sprintf(`{"jsonrpc": "2.0", "method": "mining.authorize", "params": ["user", "pass"], "id": %d}`, id))
}
