package handlers

import (
	"encoding/json"
	"github.com/inkoba/web-app/internal/model"
	"github.com/inkoba/web-app/internal/ports"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AccountHandlers struct {
	logger         *logrus.Logger
	accountService ports.IAccountService
}

func NewAccountHandlers(logger *logrus.Logger, accountService ports.IAccountService) *AccountHandlers {
	return &AccountHandlers{
		logger:         logger,
		accountService: accountService,
	}
}

func (h *AccountHandlers) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var account model.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		h.logger.Error("Unable to decode request body ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	insertResult, err := h.accountService.Create(account)
	if err != nil {
		h.logger.Error("Unable to create account ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(insertResult)
	if err != nil {
		h.logger.Error("Unable to encode response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *AccountHandlers) GetAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	accounts, err := h.accountService.GetAll()
	if err != nil {
		h.logger.Error("Unable to get accounts", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(accounts)
	if err != nil {
		h.logger.Error("Unable to encode response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *AccountHandlers) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]
	account, err := h.accountService.GetById(id)

	if err != nil {
		h.logger.Error("Unable to get accounts", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		h.logger.Error("Unable to encode response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *AccountHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	vars := mux.Vars(r)
	id := vars["id"]
	err := h.accountService.Delete(id)

	if err != nil {
		h.logger.Error("Unable to delete", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *AccountHandlers) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	vars := mux.Vars(r)
	id := vars["id"]

	var accountInput model.Account
	err := json.NewDecoder(r.Body).Decode(&accountInput)
	if err != nil {
		h.logger.Error("Unable to decode request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.accountService.Update(id, accountInput)

	if err != nil {
		h.logger.Error("Unable to update accounts", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
