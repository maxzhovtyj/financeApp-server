package repository

import (
	"context"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	Create(ctx context.Context, user models.User) error
	GetByCredentials(ctx context.Context, email, password string) (models.User, error)
}

type Repository struct {
	Users Users
}

func New(db *mongo.Database) *Repository {
	return &Repository{Users: NewUsersRepo(db)}
}
