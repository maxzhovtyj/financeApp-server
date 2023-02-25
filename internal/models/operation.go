package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Operation struct {
	Id          primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	WalletId    primitive.ObjectID   `json:"walletId" bson:"walletId"`
	Income      bool                 `json:"income" bson:"income"`
	Description string               `json:"description,omitempty" bson:"description"`
	Sum         primitive.Decimal128 `json:"sum" bson:"sum"`
	Category    Category             `json:"category,omitempty" bson:"category,omitempty"`
}
