package service

import (
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
)

type Auth interface {
	Register(newAccount *entity.Account) (*response.Account, error)
}

type AuthService struct {
	authRepo repository.Auth
}

func NewAuthService(authRepo repository.Auth) Auth {
	return &AuthService{authRepo: authRepo}
}

func (a *AuthService) Register(newAccount *entity.Account) (*response.Account, error) {
	accountResponse := &response.Account{}

	account, err := a.authRepo.Register(newAccount)
	if err != nil {
		return nil, err
	}

	accountResponse = mapper.AccountToAccountResponse(account)

	return accountResponse, nil
}
