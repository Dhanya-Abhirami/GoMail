package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Email struct {
	Id primitive.ObjectID `json:"id" bson:"_id",omitempty"`
	Sender string `json:"sender" validate:"required"`
	Receiver string `json:"receiver" validate:"required"`
	Subject string `json:"subject" validate:"required"`
	Body string `json:"body" validate:"required"`
	SentAt time.Time `json:"sentAt" validate:"required"`
}