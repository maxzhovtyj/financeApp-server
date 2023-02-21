package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Id    primitive.ObjectID
	Title string
}
