package repository

import (
	"context"
	"errors"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/pkg/db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepo struct {
	db *mongo.Collection
}

func NewUsersRepo(db *mongo.Database) *UsersRepo {
	return &UsersRepo{db: db.Collection(userCollection)}
}

func (r *UsersRepo) Create(ctx context.Context, user models.User) (primitive.ObjectID, error) {
	one, err := r.db.InsertOne(ctx, user)
	if mongodb.IsDuplicate(err) {
		return [12]byte{}, models.ErrUserAlreadyExists
	}

	return one.InsertedID.(primitive.ObjectID), err
}

func (r *UsersRepo) GetByCredentials(ctx context.Context, email, password string) (models.User, error) {
	var user models.User

	err := r.db.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, models.ErrUserNotFound
		}

		return models.User{}, err
	}

	return user, nil
}
