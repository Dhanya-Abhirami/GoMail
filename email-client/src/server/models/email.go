package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Email struct {
	Id primitive.ObjectID `json:"_id,omitempty"`
	Sender string `json:"sender" validate:"required"`
	Receiver string `json:"receiver" validate:"required"`
	Subject string `json:"subject" validate:"required"`
	Body string `json:"body" validate:"required"`
}