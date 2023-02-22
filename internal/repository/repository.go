package repository

import (
	"context"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Users interface {
	Create(ctx context.Context, user models.User) error
	GetByCredentials(ctx context.Context, email, password string) (models.User, error)
}

type Wallet interface {
	Create(ctx context.Context, wallet models.Wallet) error
}

type Operation interface {
	Create(ctx context.Context, operation models.Operation) error
}

type Transaction interface {
	StartSession(opts *options.SessionOptions) (mongo.Session, error)
}

type Repository struct {
	Users       Users
	Wallet      Wallet
	Transaction Transaction
	Operation   Operation
}

func New(db *mongo.Database) *Repository {
	return &Repository{
		Users:       NewUsersRepo(db),
		Wallet:      NewWalletRepo(db),
		Operation:   NewOperationsRepo(db),
		Transaction: NewTransactionRepo(db),
	}
}
