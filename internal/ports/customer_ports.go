package ports

import (
	"github.com/inkoba/web-app/internal/model"
	"net/http"
)

type ICustomerHandler interface {
	GetCustomers(w http.ResponseWriter, r *http.Request)
	GetOneCustomer(w http.ResponseWriter, r *http.Request)
	CreateCustomer(w http.ResponseWriter, r *http.Request)
	UpDateCustomer(w http.ResponseWriter, r *http.Request)
	DeleteCustomer(w http.ResponseWriter, r *http.Request)
}

type ICustomerService interface {
	GetCustomers() (*[]model.Customer, error)
	GetOneCustomer(id string) (*model.Customer, error)
	CreateCustomer(customer *model.Customer) (string, error)
	UpDateCustomer(id string, customer *model.Customer) error
	DeleteCustomer(id string) error
}

type ICustomerRepository interface {
	GetCustomers() (*[]model.Customer, error)
	GetOneCustomer(id string) (*model.Customer, error)
	CreateCustomer(customer *model.Customer) (string, error)
	UpDateCustomer(id string, customer *model.Customer) error
	DeleteCustomer(id string) error
}
