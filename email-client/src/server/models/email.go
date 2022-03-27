package models

// import (
// 	"errors"
// 	"strings"
// )

type Email struct {
	Sender string `json:"sender" validate:"required"`
	Receiver string `json:"receiver" validate:"required"`
	Subject string `json:"subject" validate:"required"`
	Body string `json:"subject" validate:"required"`
}