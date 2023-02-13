package repository

import "go.mongodb.org/mongo-driver/mongo"

type Users interface {
}

type Repository struct {
	Users Users
}

func New(db *mongo.Database) *Repository {
	return &Repository{Users: NewUsersRepo(db)}
}
