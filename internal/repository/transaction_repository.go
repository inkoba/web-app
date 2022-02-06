package repository

import (
	"context"
	"encoding/json"
	"github.com/inkoba/web-app/internal/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransactionRepository struct {
	logger     *logrus.Logger
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

func NewTransactionRepository(logger *logrus.Logger, cl *mongo.Client) *TransactionRepository {
	return &TransactionRepository{
		logger:     logger,
		Client:     cl,
		Database:   cl.Database("sample_analytics"),
		Collection: cl.Database("sample_analytics").Collection("transactions"),
	}
}

func (tr *TransactionRepository) GetTransaction(id primitive.ObjectID) (*model.TransactionCollection, error) {

	var transaction *model.TransactionCollection
	filter := bson.M{"_id": id}

	err := tr.Collection.FindOne(context.TODO(), filter).Decode(&transaction)
	if err != nil {
		tr.logger.Error(err)
	}
	return transaction, err
}

func (tr *TransactionRepository) GetAllTransactions() ([]*model.TransactionCollection, error) {

	var transactions []*model.TransactionCollection
	findOptions := options.Find()

	cur, err := tr.Collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		tr.logger.Error(err)
	}

	for cur.Next(context.TODO()) {
		var elem *model.TransactionCollection
		err := cur.Decode(&elem)
		if err != nil {
			tr.logger.Error(err)
		}
		transactions = append(transactions, elem)
	}

	if err := cur.Err(); err != nil {
		tr.logger.Error(err)
	}
	if err = cur.Close(context.TODO()); err != nil {
		tr.logger.Error(err)
	}

	return transactions, err
}

func (tr *TransactionRepository) PostTransaction(r []byte) error {

	var transaction *model.TransactionCollection
	err := json.Unmarshal(r, &transaction)
	if err != nil {
		tr.logger.Error(err)
	}
	_, err = tr.Collection.InsertOne(context.TODO(), transaction)
	if err != nil {
		tr.logger.Error(err)
	}
	return err
}

func (tr *TransactionRepository) UpdateTransaction(id primitive.ObjectID, r []byte) error {

	var transaction *model.TransactionCollection
	err := json.Unmarshal(r, &transaction)
	if err != nil {
		tr.logger.Error(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": &transaction}
	_, err = tr.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		tr.logger.Error(err)
	}
	return err
}

func (tr *TransactionRepository) DeleteTransaction(id primitive.ObjectID) error {

	filter := bson.M{"_id": id}

	_, err := tr.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		tr.logger.Error(err)
	}
	return err
}
