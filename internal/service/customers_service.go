package service

import (
	"github.com/inkoba/web-app/internal/model"
	"github.com/inkoba/web-app/internal/ports"
	"github.com/sirupsen/logrus"
)

type CustomerService struct {
	logger             *logrus.Logger
	customerRepository ports.ICustomerRepository
}

func NewCustomerService(logger *logrus.Logger, customerRepository ports.ICustomerRepository) *CustomerService {
	return &CustomerService{
		logger,
		customerRepository,
	}
}

func (cs *CustomerService) GetCustomers() (*[]model.Customer, error) {
	return cs.customerRepository.GetCustomers()
}

func (cs *CustomerService) GetOneCustomer(id string) (*model.Customer, error) {
	return cs.customerRepository.GetOneCustomer(id)
}

func (cs *CustomerService) CreateCustomer(customer *model.Customer) (string, error) {
	return cs.customerRepository.CreateCustomer(customer)
}

func (cs *CustomerService) UpDateCustomer(id string, customer *model.Customer) error {
	return cs.customerRepository.UpDateCustomer(id, customer)
}

func (cs *CustomerService) DeleteCustomer(id string) error {
	return cs.customerRepository.DeleteCustomer(id)
}
