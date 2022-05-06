package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Cart struct {
		ID        primitive.ObjectID `param:"id" json:"_id" bson:"_id"`
		UserId    primitive.ObjectID `json:"userId" bson:"userId"`
		Items     []CartItem         `json:"items" bson:"items"`
		CreatedAt string             `json:"createdAt" bson:"createdAt"`
		UpdatedAt string             `json:"updatedAt" bson:"updatedAt"`
	}
)