package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wallet struct {
	Id         primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	UserId     primitive.ObjectID   `json:"userId" bson:"userId"`
	Sum        primitive.Decimal128 `json:"sum" bson:"sum"`
	Operations []Operation          `json:"operations,omitempty" bson:"operations"`
}
