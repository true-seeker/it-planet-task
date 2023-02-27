package service

import (
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
)

type Account interface {
	Get(id int) (*response.Account, error)
	Search(params *filter.AccountFilterParams) (*[]response.Account, error)
	IsAlreadyExists(account *entity.Account) bool
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

func (a *AccountService) Search(params *filter.AccountFilterParams) (*[]response.Account, error) {
	var accountResponses *[]response.Account

	accounts, err := a.accountRepo.Search(params)
	if err != nil {
		return nil, err
	}

	accountResponses = mapper.AccountsToAccountResponses(accounts)

	return accountResponses, nil
}

func (a *AccountService) IsAlreadyExists(account *entity.Account) bool {
	return a.accountRepo.IsAlreadyExists(account)
}
