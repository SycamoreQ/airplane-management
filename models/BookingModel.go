package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID             primitive.ObjectID `bson:"_id"`
	Booking_ID     string             `json:"booking_id"`
	PassengerName  *string            `json:"name" validate:"required,min=2,max=100"`
	Phone          *int64             `json:"phone" validate:"required,max=10" `
	Email          *string            `json:"email" validate:"required"`
	Nationality    *string            `json:"nationality" validate:"required"`
	Airplane_ID    Airplane           `json:"airplane_id"`
	Airport_ID     Airport            `json:"airport_id"`
	Source         *string            `json:"source" validate:"required,min=2,max=100"`
	Destination    *string            `json:"destination" validate:"required,min=2,max=100"`
	Class          *string            `json:"class" validate:"required,min=2,max=100"`
	Seat_no        []int64            `json:"seat_no" validate:"required"`
	Price          *float64           `json:"price" validate:"required"`
	Terminal       Airplane           `json:"terminal" validate:"required"`
	Gate           *int64             `json:"gate" validate:"required"`
	Arrival_time   Airplane           `json:"arrival_time"`
	Departure_time Airplane           `json:"departure_time"`
	Baggage_Count  *int64             `json:"baggage_count"  validate:"required"`
	Created_At     time.Time          `json:"created_at" validate:"required"`
	Updated_At     time.Time          `json:"updated_at" validate:"required"`
}
