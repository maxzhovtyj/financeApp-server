package repository

import (
	"context"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type OperationsRepo struct {
	operations *mongo.Collection
}

func NewOperationsRepo(db *mongo.Database) *OperationsRepo {
	return &OperationsRepo{
		operations: db.Collection(operationsCollection),
	}
}

func (r *OperationsRepo) Create(ctx context.Context, operation models.Operation) error {
	if _, err := r.operations.InsertOne(ctx, operation); err != nil {
		return err
	}

	return nil
}
