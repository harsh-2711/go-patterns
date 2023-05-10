package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserType struct {
	ID        string              `json:"id" bson:"_id"`
	EmailID   *string             `json:"emailId" bson:"emailId"`
	Name      *string             `json:"name" bson:"name"`
	LastLogin *primitive.DateTime `json:"lastLogin" bson:"lastLogin"`
}
