package repository

import (
	"context"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type WalletRepo struct {
	wallets *mongo.Collection
}

func NewWalletRepo(db *mongo.Database) *WalletRepo {
	return &WalletRepo{
		wallets: db.Collection(walletsCollection),
	}
}

func (w *WalletRepo) Create(ctx context.Context, wallet models.Wallet) error {
	if _, err := w.wallets.InsertOne(ctx, wallet); err != nil {
		return err
	}

	return nil
}
