package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TransactionCollection struct {
	ID               primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	AccountID        int                `bson:"account_id" json:"account_id"`
	TransactionCount int                `bson:"transaction_count" json:"transaction_count"`
	BucketStartDate  time.Time          `bson:"bucket_start_date" json:"bucket_start_date"`
	BucketEndDate    time.Time          `bson:"bucket_end_date" json:"bucket_end_date"`
	Transactions     []Transacts        `bson:"transactions" json:"transactions"`
}

type Transacts struct {
	Date            time.Time `bson:"date" json:"date"`
	Amount          int       `bson:"amount" json:"amount"`
	TransactionCode string    `bson:"transaction_code" json:"transaction_code"`
	Symbol          string    `bson:"symbol" json:"symbol"`
	Price           string    `bson:"price" json:"price"`
	Total           string    `bson:"total" json:"total"`
}
