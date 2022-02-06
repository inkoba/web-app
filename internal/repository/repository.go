package repository

import (
	"github.com/inkoba/web-app/internal/ports"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	ports.IAccountRepository
	ports.ICustomerRepository
	ports.ITransactionRepository
}

func NewRepository(logger *logrus.Logger, db *mongo.Client) *Repository {
	return &Repository{
		IAccountRepository:     NewAccountRepository(logger, db),
		ICustomerRepository:    NewCustomerRepository(logger, db),
		ITransactionRepository: NewTransactionRepository(logger, db),
	}
}