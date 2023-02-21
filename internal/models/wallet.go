package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wallet struct {
	Id         primitive.ObjectID
	Operations []Operation
}
