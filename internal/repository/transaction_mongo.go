package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransactionRepo struct {
	client *mongo.Client
}

func NewTransactionRepo(db *mongo.Database) *TransactionRepo {
	return &TransactionRepo{client: db.Client()}
}

func (t TransactionRepo) StartSession(opts *options.SessionOptions) (mongo.Session, error) {
	session, err := t.client.StartSession(opts)
	if err != nil {
		return nil, err
	}

	return session, nil
}
