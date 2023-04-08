package AnimalTypeValidator

import (
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/validator"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

func ValidateAnimalType(animalType *entity.AnimalType) *errorHandler.HttpErr {
	if validator.IsStringEmpty(animalType.Type) {
		return errorHandler.NewHttpErr("type is empty", http.StatusBadRequest)
	}
	return nil
}
