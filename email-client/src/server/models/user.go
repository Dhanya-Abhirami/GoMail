package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id primitive.ObjectID `json:"_id,omitempty"`
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}




