package service

import (
	"github.com/inkoba/web-app/internal/model"
	"github.com/inkoba/web-app/internal/ports"
	"github.com/sirupsen/logrus"
)

type AccountService struct {
	logger *logrus.Logger
	rep    ports.IAccountRepository
}

func NewAccountService(logger *logrus.Logger, accountRepository ports.IAccountRepository) *AccountService {
	return &AccountService{
		logger: logger,
		rep:    accountRepository,
	}
}

func (s *AccountService) Create(account model.Account) (string, error) {
	return s.rep.Create(account)
}

func (s *AccountService) GetAll() ([]model.Account, error) {
	return s.rep.GetAll()
}

func (s *AccountService) GetById(accountId string) (model.Account, error) {
	return s.rep.GetById(accountId)
}

func (s *AccountService) Delete(accountId string) error {
	return s.rep.Delete(accountId)
}

func (s *AccountService) Update(accountId string, input model.Account) error {
	return s.rep.Update(accountId, input)
}
