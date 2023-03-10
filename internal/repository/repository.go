package repository

import (
	"context"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	Create(ctx context.Context, user models.User) error
	GetByCredentials(ctx context.Context, email, password string) (models.User, error)
}

type Wallet interface {
	Create(ctx context.Context, wallet models.Wallet) error
	GetWallet(ctx context.Context, walletOid primitive.ObjectID) (models.Wallet, []models.Operation, error)
	GetAllWallets(ctx context.Context, userOid primitive.ObjectID) ([]models.Wallet, error)
	NewOperation(ctx context.Context, operation models.Operation) error
}

type Repository struct {
	Users  Users
	Wallet Wallet
}

func New(db *mongo.Database) *Repository {
	return &Repository{
		Users:  NewUsersRepo(db),
		Wallet: NewWalletRepo(db),
	}
}
