package model

type Account struct {
	AccountID int32    `json:"account_id" bson:"account_id"`
	Limit     int32    `json:"limit" bson:"limit"`
	Products  []string `json:"products" bson:"products"`
}
