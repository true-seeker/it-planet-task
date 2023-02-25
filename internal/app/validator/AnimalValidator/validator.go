package AnimalValidator

import (
	"errors"
	"fmt"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

const (
	Alive = "ALIVE"
	Dead  = "DEAD"
)
const (
	Male   = "MALE"
	Female = "FEMALE"
	Other  = "OTHER"
)

func ValidateLifeStatus(lifeStatus string) *errorHandler.HttpErr {
	if lifeStatus != Alive && lifeStatus != Dead {
		return &errorHandler.HttpErr{
			Err:        errors.New(fmt.Sprintf("lifeStatus must be in [%s, %s]", Alive, Dead)),
			StatusCode: http.StatusBadRequest,
		}
	}
	return nil
}

func ValidateGender(gender string) *errorHandler.HttpErr {
	if gender != Male && gender != Female && gender != Other {
		return &errorHandler.HttpErr{
			Err:        errors.New(fmt.Sprintf("gender must be in [%s, %s, %s]", Male, Female, Other)),
			StatusCode: http.StatusBadRequest,
		}
	}
	return nil
}
