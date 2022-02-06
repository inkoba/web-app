package repository

import (
	"context"
	"fmt"
	"github.com/inkoba/web-app/internal/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	database   = "sample_analytics"
	collection = "customers"
)

var Ctx = context.Background()

type CustomerRepository struct {
	logger     *logrus.Logger
	collection *mongo.Collection
}

func NewCustomerRepository(logger *logrus.Logger, client *mongo.Client) *CustomerRepository {
	return &CustomerRepository{
		logger,
		client.Database(database).Collection(collection),
	}
}

func (c *CustomerRepository) GetCustomers() (*[]model.Customer, error) {
	var customers []model.Customer
	var cursor, err = c.collection.Find(Ctx, bson.D{})

	if err != nil {
		defer func(cursor *mongo.Cursor, ctx context.Context) {
			var err = cursor.Close(ctx)
			if err != nil {
				c.logger.Error("Error in find all documents in database mongodb", err)
			}
		}(cursor, Ctx)
		return &customers, err
	}

	var customer model.Customer
	for cursor.Next(Ctx) {
		err := cursor.Decode(&customer)
		if err != nil {
			return &customers, err
		}
		customers = append(customers, customer)
	}
	return &customers, nil
}

func (c *CustomerRepository) GetOneCustomer(id string) (*model.Customer, error) {
	var customer model.Customer
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.logger.Error("Error in repository give objectId", err)
	}

	err = c.collection.FindOne(Ctx, bson.M{"_id": objectId}).Decode(&customer)
	if err != nil {
		c.logger.Error("Error in find one documents in database mongodb", err)
	}

	return &customer, err
}

func (c *CustomerRepository) CreateCustomer(customer *model.Customer) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := c.collection.InsertOne(ctx, customer)
	if err != nil {
		return "0", err
	}
	id := fmt.Sprintf("%v", res.InsertedID)
	return id, err
}

func (c *CustomerRepository) UpDateCustomer(id string, customer *model.Customer) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", customer}}
	_, err = c.collection.UpdateOne(Ctx, filter, update)
	if err != nil {
		c.logger.Error("Error in find one documents in database mongodb: ", err)
	}
	return err
}

func (c *CustomerRepository) DeleteCustomer(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = c.collection.DeleteOne(Ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}
