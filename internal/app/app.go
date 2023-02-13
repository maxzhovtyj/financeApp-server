package app

import (
	"fmt"
	"github.com/maxzhovtyj/financeApp-server/internal/config"
	delivery "github.com/maxzhovtyj/financeApp-server/internal/delivery/http"
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
	"github.com/maxzhovtyj/financeApp-server/internal/service"
	"github.com/maxzhovtyj/financeApp-server/pkg/db/mongodb"
	"log"
)

func Run() {
	// Init configs
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.Mongo)

	// Init mongo client
	dbClient := mongodb.New()

	// Init repository
	repo := repository.New(dbClient.Database(""))

	// Init services
	s := service.New(repo)

	// Init handlers
	h := delivery.New(s)

	echo := h.Init()

	echo.Server.WriteTimeout = cfg.HTTP.WriteTimeout
	echo.Server.ReadTimeout = cfg.HTTP.ReadTimeout
	echo.Server.MaxHeaderBytes = cfg.HTTP.MaxHeaderMegabytes << 20

	// Run application
	if err = echo.Start(":" + cfg.HTTP.Port); err != nil {
		log.Fatal(err)
	}
}
