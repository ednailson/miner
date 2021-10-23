package main

import (
	"github.com/ednailson/miner/internal/domain/usecase"
	"github.com/ednailson/miner/pkg/database"
	"github.com/ednailson/miner/pkg/log"
	"github.com/ednailson/miner/pkg/server"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	log.Setup()

	db, err := database.Setup()
	if err != nil {
		zap.S().Fatalf("database: failed to initialize database, error: %s", err.Error())
	}

	server.Setup(usecase.NewMiner(db))
}
