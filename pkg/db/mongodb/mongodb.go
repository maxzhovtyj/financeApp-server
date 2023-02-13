package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func New() *mongo.Client {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	dsn := fmt.Sprintf("mongodb://localhost:2717")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn).SetAuth(options.Credential{
		Username: "admin",
		Password: "30042003",
	}))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
