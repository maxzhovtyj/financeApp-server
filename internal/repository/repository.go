package repository

import (
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	SignUp(user models.User) (primitive.ObjectID, error)
}

type Repository struct {
	Users Users
}

func New(db *mongo.Database) *Repository {
	return &Repository{Users: NewUsersRepo(db)}
}
