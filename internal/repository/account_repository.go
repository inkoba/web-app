package repository

import (
	"context"
	"fmt"
	"time"
	"github.com/sirupsen/logrus"
	"github.com/inkoba/web-app/internal/model"
	"github.com/inkoba/web-app/internal/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountRepository struct {
	logger     *logrus.Logger
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

var _ ports.IAccountRepository = (*AccountRepository)(nil)

func NewAccountRepository(logger *logrus.Logger, client *mongo.Client) *AccountRepository {
	return &AccountRepository{
		logger:     logger,
		client:     client,
		database:   client.Database("sample_analytics"),
		collection: client.Database("sample_analytics").Collection("accounts"),
	}
}

func (rep *AccountRepository) Create(account model.Account) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := rep.collection.InsertOne(ctx, account)
	if err != nil {
		return "0", err
	}
	id := fmt.Sprintf("%v", res.InsertedID)
	return id, err
}

func (rep *AccountRepository) GetAll() ([]model.Account, error) {
	var accounts []model.Account

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := rep.collection.Find(ctx, bson.D{})

	if err != nil {
		defer func(cursor *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {
				rep.logger.Error("Get all request to db ", err)
			}
		}(cursor, ctx)
		return accounts, err
	}

	for cursor.Next(context.TODO()) {
		var account model.Account
		err := cursor.Decode(&account)
		if err != nil {
			return accounts, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (rep *AccountRepository) GetById(accountId string) (model.Account, error) {
	var account model.Account
	objectId, err := primitive.ObjectIDFromHex(accountId)
	if err != nil {
		return account, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = rep.collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&account)
	if err != nil {
		return account, err
	}
	return account, nil
}

func (rep *AccountRepository) Delete(accountId string) error {
	objectId, err := primitive.ObjectIDFromHex(accountId)
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = rep.collection.DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}

func (rep *AccountRepository) Update(accountId string, input model.Account) error {
	objectId, err := primitive.ObjectIDFromHex(accountId)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", input}}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = rep.collection.UpdateOne(ctx, filter, update)
	return err
}
