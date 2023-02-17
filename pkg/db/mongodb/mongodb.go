package mongodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/maxzhovtyj/financeApp-server/internal/config"
	"github.com/maxzhovtyj/financeApp-server/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func New(cfg config.MongoConfig) *mongo.Client {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	dsn := fmt.Sprintf(cfg.URI)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn).SetAuth(options.Credential{
		Username: cfg.User,
		Password: cfg.Password,
	}))
	if err != nil {
		logger.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Fatal(err)
	}

	return client
}

func IsDuplicate(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}

	return false
}
