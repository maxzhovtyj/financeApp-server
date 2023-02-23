package app

import (
	"github.com/maxzhovtyj/financeApp-server/internal/config"
	delivery "github.com/maxzhovtyj/financeApp-server/internal/delivery/http"
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
	"github.com/maxzhovtyj/financeApp-server/internal/server"
	"github.com/maxzhovtyj/financeApp-server/internal/service"
	"github.com/maxzhovtyj/financeApp-server/pkg/auth"
	"github.com/maxzhovtyj/financeApp-server/pkg/db/mongodb"
	"github.com/maxzhovtyj/financeApp-server/pkg/hash"
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

	// Init token manager client
	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		logger.Fatal("failed to import token manager")
	}

	// Init password hashing package
	hashing := hash.NewSHA1Hashing(cfg.Auth.PasswordSalt)

	// Init repository, services and handlers
	repo := repository.New(dbClient.Database(cfg.Mongo.Database))
	s := service.New(repo, tokenManager, cfg.Auth.JWT.AccessTokenTTL, cfg.Auth.JWT.RefreshTokenTTL, hashing)
	h := delivery.New(s, tokenManager)

	// Init and run server
	srv := server.NewServer(cfg, h.Init())
	if err = srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
