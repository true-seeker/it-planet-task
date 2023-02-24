package service

import (
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
)

type Account interface {
	Get(id int) (*response.Account, error)
}

type AccountService struct {
	accountRepo repository.Account
}

func NewAccountService(accountRepo repository.Account) Account {
	return &AccountService{accountRepo: accountRepo}
}

func (a *AccountService) Get(id int) (*response.Account, error) {
	accountResponse := &response.Account{}

	account, err := a.accountRepo.Get(id)
	if err != nil {
		return nil, err
	}

	accountResponse.Id = account.Id
	accountResponse.Email = account.Email
	accountResponse.FirstName = account.FirstName
	accountResponse.LastName = account.LastName

	return accountResponse, nil
}
