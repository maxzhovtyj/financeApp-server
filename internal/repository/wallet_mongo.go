package repository

import "go.mongodb.org/mongo-driver/mongo"

type WalletRepo struct {
	wallets *mongo.Collection
}

func NewWalletRepo(db *mongo.Database) *WalletRepo {
	return &WalletRepo{wallets: db.Collection(walletsCollection)}
}
