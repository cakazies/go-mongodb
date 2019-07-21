package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v8"
)

type (
	Student struct {
		ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
		Firstname string             `json:"firstname" bson:"firstname" validate:"required"`
		Lastname  string             `json:"lastname" bson:"lastname" validate:"required"`
	}

	Response struct {
		Data    []Student `json:"data,omitempty"`
		Message string    `json:"message,omitempty"`
	}

	CheckValidator struct {
		validator *validator.Validate
	}
)

func (cv *CheckValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Resp(data []Student, message string) interface{} {
	var resp Response
	resp.Data = data
	resp.Message = message
	return resp
}
