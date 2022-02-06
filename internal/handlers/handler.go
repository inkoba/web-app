package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/inkoba/web-app/internal/ports"
	"github.com/inkoba/web-app/internal/service"
)

type Handler struct {
	ports.IAccountHandlers
	ports.ICustomerHandler
	ports.ITransactionHandlers
}

func NewHandler(logger *logrus.Logger, s *service.Service) *Handler {
	return &Handler{
		IAccountHandlers:     NewAccountHandlers(logger, s.IAccountService),
		ICustomerHandler:     NewCustomerHandlers(logger, s.ICustomerService),
		ITransactionHandlers: NewTransactionHandlers(logger, s.ITransactionService),
	}
}
