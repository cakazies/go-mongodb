package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}
type Response struct {
	Data    []Student `json:"data,omitempty"`
	Message string    `json:"message,omitempty"`
}

func Resp(data []Student, message string) interface{} {
	var resp Response
	resp.Data = data
	resp.Message = message
	return resp
}
