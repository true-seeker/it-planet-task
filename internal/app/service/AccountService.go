package service

import (
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
	"net/url"
)

type Account interface {
	Get(id int) (*response.Account, error)
	Search(query url.Values) (*[]response.Account, error)
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

	accountResponse = mapper.AccountToAccountResponse(account)

	return accountResponse, nil
}

func (a *AccountService) Search(query url.Values) (*[]response.Account, error) {
	var accountResponses *[]response.Account

	accounts, err := a.accountRepo.Search(query)
	if err != nil {
		return nil, err
	}

	accountResponses = mapper.AccountsToAccountResponses(accounts)

	return accountResponses, nil
}
