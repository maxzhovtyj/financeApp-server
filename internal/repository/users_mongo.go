package repository

import (
	"context"
	"errors"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/pkg/db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepo struct {
	users *mongo.Collection
}

func NewUsersRepo(db *mongo.Database) *UsersRepo {
	return &UsersRepo{users: db.Collection(usersCollection)}
}

func (r *UsersRepo) Create(ctx context.Context, user models.User) error {
	_, err := r.users.InsertOne(ctx, user)
	if mongodb.IsDuplicate(err) {
		return models.ErrUserAlreadyExists
	}

	return err
}

func (r *UsersRepo) GetByCredentials(ctx context.Context, email, password string) (models.User, error) {
	var user models.User

	err := r.users.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, models.ErrUserNotFound
		}

		return models.User{}, err
	}

	return user, nil
}
