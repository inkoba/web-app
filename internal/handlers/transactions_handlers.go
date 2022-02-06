package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/inkoba/web-app/internal/ports"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
)

type TransactionHandlers struct {
	transactionService ports.ITransactionService
	logger             *logrus.Logger
}

func NewTransactionHandlers(logger *logrus.Logger, transactionService ports.ITransactionService) *TransactionHandlers {
	return &TransactionHandlers{
		transactionService: transactionService,
		logger:             logger,
	}
}

func (db *TransactionHandlers) GetTransactionEndpoint(w http.ResponseWriter, r *http.Request) {

	db.logger.Info("GetTransaction endpoint start")
	vars := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(vars["id"])

	tr, err := db.transactionService.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			db.logger.Error(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(tr)
	if err != nil {
		db.logger.Error(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		db.logger.Error(err)
	}

	db.logger.Info("GetTransaction endpoint end")

}

func (db *TransactionHandlers) GetAllTransactionsEndpoint(w http.ResponseWriter, _ *http.Request) {

	db.logger.Info("GetAllTransactions endpoint start")
	transactions, err := db.transactionService.GetAllTransactions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(`{"message": ` + err.Error() + `"}`))
		if err != nil {
			db.logger.Error(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(transactions)
	if err != nil {
		db.logger.Error(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		db.logger.Error(err)
	}
	db.logger.Info("GetAllTransactions endpoint end")

}

func (db *TransactionHandlers) PostTransactionEndpoint(w http.ResponseWriter, r *http.Request) {

	db.logger.Info("PostTransaction endpoint start")

	postBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		db.logger.Error(err)
	}

	err = db.transactionService.PostTransaction(postBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			db.logger.Error(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := "Successfully added"
	_, err = w.Write([]byte(response))
	if err != nil {
		db.logger.Error(err)
	}

	db.logger.Info("PostTransaction endpoint end")

}

func (db *TransactionHandlers) UpdateTransactionEndpoint(w http.ResponseWriter, r *http.Request) {

	db.logger.Info("UpdateTransactionEndpoint endpoint start")

	vars := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(vars["id"])

	putBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		db.logger.Error(err)
	}

	err = db.transactionService.UpdateTransaction(id, putBody)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			db.logger.Error(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := "Successfully updated"
	_, err = w.Write([]byte(response))
	if err != nil {
		db.logger.Error(err)
	}

	db.logger.Info("UpdateTransactionEndpoint endpoint end")

}

func (db *TransactionHandlers) DeleteTransactionEndpoint(w http.ResponseWriter, r *http.Request) {

	db.logger.Info("DeleteTransactionEndpoint endpoint start")

	vars := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(vars["id"])

	err := db.transactionService.DeleteTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			db.logger.Error(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := "Successfully deleted"
	_, err = w.Write([]byte(response))
	if err != nil {
		db.logger.Error(err)
	}

	db.logger.Info("DeleteTransactionEndpoint endpoint end")

}
