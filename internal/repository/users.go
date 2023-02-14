package repository

import (
	"context"
	"fmt"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepo struct {
	db *mongo.Collection
}

func NewUsersRepo(db *mongo.Database) *UsersRepo {
	return &UsersRepo{db: db.Collection(userCollection)}
}

func (r *UsersRepo) SignUp(user models.User) (primitive.ObjectID, error) {
	one, err := r.db.InsertOne(context.Background(), user)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("failed to insert user: %v", err)
	}

	return one.InsertedID.(primitive.ObjectID), nil
}
