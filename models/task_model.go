package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Description string             `json:"description,omitempty" validate:"required"`
	IsCompleted *bool              `json:"isCompleted,omitempty" validate:"required"`
}
