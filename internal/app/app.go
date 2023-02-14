package app

import (
	"github.com/maxzhovtyj/financeApp-server/internal/config"
	delivery "github.com/maxzhovtyj/financeApp-server/internal/delivery/http"
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
	"github.com/maxzhovtyj/financeApp-server/internal/server"
	"github.com/maxzhovtyj/financeApp-server/internal/service"
	"github.com/maxzhovtyj/financeApp-server/pkg/db/mongodb"
	"github.com/maxzhovtyj/financeApp-server/pkg/logger"
)

func Run() {
	// Init configs
	cfg, err := config.Init()
	if err != nil {
		logger.Fatal(err)
	}

	// Init mongo client
	dbClient := mongodb.New(cfg.Mongo)

	// Init repository, services and handlers
	repo := repository.New(dbClient.Database(cfg.Mongo.Database))
	s := service.New(repo)
	h := delivery.New(s)

	srv := server.NewServer(cfg, h.Init())
	if err = srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
