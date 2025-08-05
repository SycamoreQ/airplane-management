package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Airplane struct {
	ID             primitive.ObjectID `bson:"_id"`
	Airplane_ID    string             `json:"airplane_id"`
	Airline        *string            `json:"airline" validate:"required,min=2,max=100"`
	Model          *string            `json:"model" validate:"required,min=2,max=100"`
	Capacity       *int64             `json:"capacity" validate:"required,max=300"`
	Top_speed      *float64           `json:"top_speed" validate:"required"`
	Destination    *string            `json:"destination" validate:"required"`
	Terminal       *int64             `json:"terminal" validate:"required"`
	Runway         *int64             `json:"runway" validate:"required"`
	Arrival_time   time.Time          `json:"arrival_time"`
	Departure_time time.Time          `json:"departure_time"`
	Created_At     time.Time          `json:"created_at" validate:"required"`
	Updated_At     time.Time          `json:"updated_at" validate:"required"`
}
