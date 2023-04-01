package AccountValidator

import (
	"errors"
	"fmt"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/validator"
	"it-planet-task/pkg/errorHandler"
	"net/http"
	"net/mail"
)

func ValidateAccount(account *entity.Account) *errorHandler.HttpErr {
	if validator.IsStringEmpty(account.FirstName) {
		return &errorHandler.HttpErr{
			Err:        errors.New("firstName is empty"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if validator.IsStringEmpty(account.LastName) {
		return &errorHandler.HttpErr{
			Err:        errors.New("lastName is empty"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if validator.IsStringEmpty(account.Email) {
		return &errorHandler.HttpErr{
			Err:        errors.New("email is empty"),
			StatusCode: http.StatusBadRequest,
		}
	}
	_, err := mail.ParseAddress(account.Email)
	if err != nil {
		return &errorHandler.HttpErr{
			Err:        errors.New("email is invalid"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if validator.IsStringEmpty(account.Password) {
		return &errorHandler.HttpErr{
			Err:        errors.New("password is empty"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if account.Role != entity.AdminRole && account.Role != entity.ChipperRole && account.Role != entity.UserRole {
		return &errorHandler.HttpErr{
			Err:        errors.New(fmt.Sprintf("role must be in [%s, %s, %s]", entity.AdminRole, entity.ChipperRole, entity.UserRole)),
			StatusCode: http.StatusBadRequest,
		}
	}

	return nil
}
