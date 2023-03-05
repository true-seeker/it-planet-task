package AnimalTypeValidator

import (
	"errors"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/validator"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

func ValidateAnimalType(animalType *entity.AnimalType) *errorHandler.HttpErr {
	if validator.IsStringEmpty(animalType.Type) {
		return &errorHandler.HttpErr{
			Err:        errors.New("type is empty"),
			StatusCode: http.StatusBadRequest,
		}
	}
	return nil
}
