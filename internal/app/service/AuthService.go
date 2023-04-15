package service

import (
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
	"it-planet-task/pkg/auth"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

type Auth interface {
	Register(newAccount *entity.Account) (*response.Account, error)
	Login(account *entity.Account) (*entity.AuthToken, *errorHandler.HttpErr)
}

type AuthService struct {
	authRepo repository.Auth
}

func NewAuthService(authRepo repository.Auth) Auth {
	return &AuthService{authRepo: authRepo}
}

func (a *AuthService) Register(newAccount *entity.Account) (*response.Account, error) {
	accountResponse := &response.Account{}

	newAccount.Role = entity.UserRole
	account, err := a.authRepo.Register(newAccount)
	if err != nil {
		return nil, err
	}

	accountResponse = mapper.AccountToAccountResponse(account)

	return accountResponse, nil
}

func (a *AuthService) Login(account *entity.Account) (*entity.AuthToken, *errorHandler.HttpErr) {
	authToken := &entity.AuthToken{
		AccountId: account.Id,
		Token:     auth.GenerateSecureToken(128),
	}

	authToken, err := a.authRepo.SaveToken(authToken)
	if err != nil {
		return nil, &errorHandler.HttpErr{
			Err:        err,
			StatusCode: http.StatusBadRequest,
		}
	}

	return authToken, nil

}
