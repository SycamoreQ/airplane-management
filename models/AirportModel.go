package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Airport struct {
	ID         primitive.ObjectID `bson:"_id"`
	Airport_ID string             `json:"airport_id"`
	Name       *string            `json:"name" validate:"required,min=2,max=100"`
	City       *string            `json:"city" validate:"required,min=2,max=100"`
	Country    *string            `json:"country" validate:"required,min=2,max=100"`
	Is_active  *bool              `json:"is_active" validate:"required"`
	Created_At time.Time          `json:"created_at" validate:"required"`
	Updated_At time.Time          `json:"updated_at" validate:"required"`
}
