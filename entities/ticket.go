package entities

import "time"

type Ticket struct {
	CreateAt       time.Time `json:"CreateAt" bson:"create_at"`
	ExpirationDate time.Time `json:"ExpirationDate" bson:"expiration_date"`
}
