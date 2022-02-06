package service

import (
	"github.com/inkoba/web-app/internal/model"
	"github.com/inkoba/web-app/internal/ports"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionService struct {
	trRep  ports.ITransactionRepository
	logger *logrus.Logger
}

func NewTransactionService(logger *logrus.Logger, repo ports.ITransactionRepository) *TransactionService {
	return &TransactionService{
		trRep:  repo,
		logger: logger,
	}
}

func (s *TransactionService) PostTransaction(r []byte) error {
	return s.trRep.PostTransaction(r)
}

func (s *TransactionService) GetAllTransactions() ([]*model.TransactionCollection, error) {
	return s.trRep.GetAllTransactions()
}

func (s *TransactionService) GetTransaction(id primitive.ObjectID) (*model.TransactionCollection, error) {

	return s.trRep.GetTransaction(id)
}

func (s *TransactionService) DeleteTransaction(id primitive.ObjectID) error {
	return s.trRep.DeleteTransaction(id)
}

func (s *TransactionService) UpdateTransaction(id primitive.ObjectID, r []byte) error {
	return s.trRep.UpdateTransaction(id, r)
}
