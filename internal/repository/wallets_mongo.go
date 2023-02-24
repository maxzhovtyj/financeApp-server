package repository

import (
	"context"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WalletRepo struct {
	wallets    *mongo.Collection
	operations *mongo.Collection
}

func NewWalletRepo(db *mongo.Database) *WalletRepo {
	return &WalletRepo{
		wallets:    db.Collection(walletsCollection),
		operations: db.Collection(operationsCollection),
	}
}

func (r *WalletRepo) Create(ctx context.Context, wallet models.Wallet) error {
	if _, err := r.wallets.InsertOne(ctx, wallet); err != nil {
		return err
	}

	return nil
}

func (r *WalletRepo) NewOperation(ctx context.Context, operation models.Operation) error {
	session, err := r.wallets.Database().Client().StartSession(&options.SessionOptions{})
	if err != nil {
		return err
	}

	if err = session.StartTransaction(&options.TransactionOptions{}); err != nil {
		return err
	}

	if _, err = r.operations.InsertOne(ctx, operation); err != nil {
		if err = session.AbortTransaction(ctx); err != nil {
			return err
		}

		return err
	}

	if operation.Income == true {
		update := r.wallets.FindOneAndUpdate(
			ctx,
			bson.M{"_id": operation.WalletId}, bson.M{"$inc": bson.M{"sum": operation.Sum}},
		)
		if update.Err() != nil {
			if err = session.AbortTransaction(ctx); err != nil {
				return err
			}
		}
	}

	if err = session.CommitTransaction(ctx); err != nil {
		return err
	}

	session.EndSession(ctx)

	return nil
}