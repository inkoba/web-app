package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/inkoba/web-app/internal/model"
	"github.com/inkoba/web-app/internal/ports"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	headersType, headersValue = "Content-Type", "application/json"
	statusHTTP                = http.StatusOK
)

type CustomerHandler struct {
	logger          *logrus.Logger
	customerService ports.ICustomerService
}

func NewCustomerHandlers(logger *logrus.Logger, CustomerService ports.ICustomerService) *CustomerHandler {
	return &CustomerHandler{
		logger,
		CustomerService,
	}
}

func (ch *CustomerHandler) GetCustomers(w http.ResponseWriter, _ *http.Request) {
	ch.logger.Info("GetCustomers endpoint start")
	w.Header().Set(headersType, headersValue)
	w.WriteHeader(statusHTTP)

	customers, err := ch.customerService.GetCustomers()

	if err != nil {
		ch.logger.Error("Unable to get customers", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(customers)
	if err != nil {
		ch.logger.Error("Unable to encode response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ch.logger.Info("GetCustomers endpoint end")
}

func (ch *CustomerHandler) GetOneCustomer(w http.ResponseWriter, r *http.Request) {
	ch.logger.Info("GetOneCustomer endpoint start")
	w.Header().Set(headersType, headersValue)
	w.WriteHeader(statusHTTP)

	vars := mux.Vars(r)
	id := vars["id"]

	customer, err := ch.customerService.GetOneCustomer(id)
	if err != nil {
		ch.logger.Error("Unable to get customer", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		ch.logger.Error("Unable to encode response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ch.logger.Info("GetOneCustomer endpoint end")
}

func (ch *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	ch.logger.Info("CreateCustomer endpoint start")
	w.Header().Set(headersType, headersValue)
	w.WriteHeader(statusHTTP)

	var customer model.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		ch.logger.Error("Unable to decode request body ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := ch.customerService.CreateCustomer(&customer)
	if err != nil {
		ch.logger.Error("Unable to create customer ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		ch.logger.Error("Unable to encode response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ch.logger.Info("CreateCustomer endpoint end")
}

func (ch *CustomerHandler) UpDateCustomer(w http.ResponseWriter, r *http.Request) {
	ch.logger.Info("UpDateCustomer endpoint start")
	w.Header().Set(headersType, headersValue)
	w.WriteHeader(statusHTTP)

	vars := mux.Vars(r)
	id := vars["id"]

	var customer model.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		ch.logger.Error("Unable to decode request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = ch.customerService.UpDateCustomer(id, &customer)
	if err != nil {
		ch.logger.Error("Unable to update customers", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ch.logger.Info("UpDateCustomer endpoint end")

}

func (ch *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	ch.logger.Info("DeleteCustomer endpoint start")
	w.Header().Set(headersType, headersValue)
	w.WriteHeader(statusHTTP)

	vars := mux.Vars(r)
	id := vars["id"]
	err := ch.customerService.DeleteCustomer(id)

	if err != nil {
		ch.logger.Error("Unable to delete", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ch.logger.Info("DeleteCustomer endpoint end")
}
