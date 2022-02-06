package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const timeout = 10

type AuthMongo struct {
	client *mongo.Client
}

func (m *AuthMongo) GetClient() *mongo.Client {
	return m.client
}

func NewAuth(uri string) *AuthMongo {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Unable connect to database ", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)

	if err = client.Connect(ctx); err != nil {
		log.Fatal("Unable connect to database ", err)
	}
	return &AuthMongo{client: client}
}
