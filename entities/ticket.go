package entities

type Ticket struct {
	CreateAt       int64 `json:"CreateAt" bson:"create_at"`
	ExpirationDate int64 `json:"ExpirationDate" bson:"expiration_date"`
}
