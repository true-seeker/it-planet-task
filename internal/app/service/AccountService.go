package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

type Account interface {
	Get(id int) (*response.Account, *errorHandler.HttpErr)
	GetByEmail(account *entity.Account) (*response.Account, error)
	Update(account *entity.Account) (*response.Account, error)
	Search(params *filter.AccountFilterParams) (*[]response.Account, error)
	IsAlreadyExists(account *entity.Account) bool
	GetByCreds(account *entity.Account) *entity.Account
	Delete(id int) error
	Create(account *entity.Account) (*response.Account, error)
}

type AccountService struct {
	accountRepo repository.Account
}

func NewAccountService(accountRepo repository.Account) Account {
	return &AccountService{accountRepo: accountRepo}
}

func (a *AccountService) Get(id int) (*response.Account, *errorHandler.HttpErr) {
	accountResponse := &response.Account{}

	account, err := a.accountRepo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorHandler.HttpErr{
				Err:        errors.New(fmt.Sprintf("Account with id %d does not exists", id)),
				StatusCode: http.StatusNotFound,
			}
		} else {
			return nil, &errorHandler.HttpErr{
				Err:        err,
				StatusCode: http.StatusBadRequest,
			}
		}
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
	return a.accountRepo.GetByEmail(account).Id != 0
}

func (a *AccountService) GetByCreds(account *entity.Account) *entity.Account {
	return a.accountRepo.GetByCreds(account)
}

func (a *AccountService) Update(account *entity.Account) (*response.Account, error) {
	accountResponse := &response.Account{}

	account, err := a.accountRepo.Update(account)
	if err != nil {
		return nil, err
	}

	accountResponse = mapper.AccountToAccountResponse(account)

	return accountResponse, nil
}

func (a *AccountService) GetByEmail(account *entity.Account) (*response.Account, error) {
	accountResponse := &response.Account{}
	account = a.accountRepo.GetByEmail(account)
	accountResponse = mapper.AccountToAccountResponse(account)
	return accountResponse, nil
}

func (a *AccountService) Delete(id int) error {
	return a.accountRepo.Delete(id)
}

func (a *AccountService) Create(account *entity.Account) (*response.Account, error) {
	accountResponse := &response.Account{}

	account, err := a.accountRepo.Create(account)
	if err != nil {
		return nil, err
	}

	accountResponse = mapper.AccountToAccountResponse(account)

	return accountResponse, nil
}
