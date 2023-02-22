package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Operation struct {
	Income      bool
	WalletId    primitive.ObjectID
	Category    Category
	Description string
	Sum         float64
}
