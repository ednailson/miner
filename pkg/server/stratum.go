package server

import (
	"github.com/ednailson/miner/internal/domain/usecase"
	"github.com/ednailson/miner/internal/infra/server"
	"go.uber.org/zap"
	"os"
)

func Setup(uc usecase.Miner) {
	addr := os.Getenv("SERVER_ADDRESS")

	s := server.NewServer(addr, uc)

	if err := s.Start(); err != nil {
		zap.S().Errorf("server: failed to start server, error: %s", err.Error())
	}
	defer s.Close()
}
