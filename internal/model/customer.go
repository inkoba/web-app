package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	Name       string             `json:"name" bson:"name"`
	Username   string             `json:"username" bson:"username"`
	IdCustomer primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email      string             `json:"email" bson:"email"`
	Address    string             `json:"address" bson:"address"`
}
