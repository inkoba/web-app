package service

import (
	"github.com/inkoba/web-app/internal/ports"
	"github.com/inkoba/web-app/internal/repository"
	"github.com/sirupsen/logrus"
)

type Service struct {
	ports.IAccountService
	ports.ICustomerService
	ports.ITransactionService
}

func NewService(logger *logrus.Logger, r *repository.Repository) *Service {
	return &Service{
		IAccountService:     NewAccountService(logger, r.IAccountRepository),
		ICustomerService:    NewCustomerService(logger, r.ICustomerRepository),
		ITransactionService: NewTransactionService(logger, r.ITransactionRepository),
	}
}
