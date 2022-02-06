package ports

import (
	"github.com/inkoba/web-app/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type ITransactionHandlers interface {
	GetTransactionEndpoint(w http.ResponseWriter, r *http.Request)
	GetAllTransactionsEndpoint(w http.ResponseWriter, r *http.Request)
	PostTransactionEndpoint(w http.ResponseWriter, r *http.Request)
	UpdateTransactionEndpoint(w http.ResponseWriter, r *http.Request)
	DeleteTransactionEndpoint(w http.ResponseWriter, r *http.Request)
}

type ITransactionRepository interface {
	GetTransaction(id primitive.ObjectID) (*model.TransactionCollection, error)
	GetAllTransactions() ([]*model.TransactionCollection, error)
	PostTransaction(r []byte) error
	UpdateTransaction(id primitive.ObjectID, r []byte) error
	DeleteTransaction(id primitive.ObjectID) error
}

type ITransactionService interface {
	GetTransaction(id primitive.ObjectID) (*model.TransactionCollection, error)
	GetAllTransactions() ([]*model.TransactionCollection, error)
	PostTransaction(r []byte) error
	UpdateTransaction(id primitive.ObjectID, r []byte) error
	DeleteTransaction(id primitive.ObjectID) error
}
