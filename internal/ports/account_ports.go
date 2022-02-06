package ports

import (
	"github.com/inkoba/web-app/internal/model"
	"net/http"
)

type IAccountService interface {
	Create(account model.Account) (string, error)
	GetAll() ([]model.Account, error)
	GetById(accountId string) (model.Account, error)
	Delete(accountId string) error
	Update(accountId string, input model.Account) error
}

type IAccountRepository interface {
	Create(account model.Account) (string, error)
	GetAll() ([]model.Account, error)
	GetById(accountId string) (model.Account, error)
	Delete(accountId string) error
	Update(accountId string, input model.Account) error
}

type IAccountHandlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}
