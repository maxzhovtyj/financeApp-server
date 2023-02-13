package app

import (
	"github.com/maxzhovtyj/financeApp-server/internal/config"
	delivery "github.com/maxzhovtyj/financeApp-server/internal/delivery/http"
	"github.com/maxzhovtyj/financeApp-server/internal/repository"
	"github.com/maxzhovtyj/financeApp-server/internal/service"
	"github.com/maxzhovtyj/financeApp-server/pkg/db/mongodb"
)

type App struct {
}

type Database struct {
}

func Run() {
	// Init configs
	config.Init()

	// Init mongo client
	dbClient := mongodb.New()

	// Init repository
	repo := repository.New(dbClient.Database(""))

	// Init services
	s := service.New(repo)

	// Init handlers
	h := delivery.New(s)
	err := h.Init().Start(":8080")
	if err != nil {
		return
	}
	// Run application

}
