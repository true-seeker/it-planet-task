package AccountValidator

import (
	"fmt"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/validator"
	"it-planet-task/pkg/errorHandler"
	"net/http"
	"net/mail"
)

func ValidateAccountRegistration(account *entity.Account) *errorHandler.HttpErr {
	if validator.IsStringEmpty(account.FirstName) {
		return errorHandler.NewHttpErr("firstName is empty", http.StatusBadRequest)
	}

	if validator.IsStringEmpty(account.LastName) {
		return errorHandler.NewHttpErr("lastName is empty", http.StatusBadRequest)
	}

	if validator.IsStringEmpty(account.Email) {
		return errorHandler.NewHttpErr("email is empty", http.StatusBadRequest)
	}
	_, err := mail.ParseAddress(account.Email)
	if err != nil {
		return errorHandler.NewHttpErr("email is invalid", http.StatusBadRequest)
	}

	if validator.IsStringEmpty(account.Password) {
		return errorHandler.NewHttpErr("password is empty", http.StatusBadRequest)
	}
	return nil
}

func ValidateAccount(account *entity.Account) *errorHandler.HttpErr {
	httpErr := ValidateAccountRegistration(account)
	if httpErr != nil {
		return httpErr
	}

	if account.Role != entity.AdminRole && account.Role != entity.ChipperRole && account.Role != entity.UserRole {
		return errorHandler.NewHttpErr(fmt.Sprintf("role must be in [%s, %s, %s]", entity.AdminRole, entity.ChipperRole, entity.UserRole), http.StatusBadRequest)
	}

	return nil
}
