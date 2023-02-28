package AccountValidator

import (
	"errors"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/pkg/errorHandler"
	"net/http"
	"net/mail"
	"strings"
)

func ValidateAccount(account *entity.Account) *errorHandler.HttpErr {
	if strings.Trim(account.FirstName, " \t\n") == "" {
		return &errorHandler.HttpErr{
			Err:        errors.New("firstName is empty"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if strings.Trim(account.LastName, " \t\n") == "" {
		return &errorHandler.HttpErr{
			Err:        errors.New("lastName is empty"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if strings.Trim(account.Email, " \t\n") == "" {
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

	if strings.Trim(account.Password, " \t\n") == "" {
		return &errorHandler.HttpErr{
			Err:        errors.New("password is empty"),
			StatusCode: http.StatusBadRequest,
		}
	}

	return nil
}
