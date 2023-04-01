package AreaValidator

import (
	"errors"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AreaPointValidator"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

func ValidateArea(area *entity.Area) *errorHandler.HttpErr {
	// TODO совместить сообщения об ошибках
	if validator.IsStringEmpty(area.Name) {
		return &errorHandler.HttpErr{
			Err:        errors.New("name is empty"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if area.AreaPoints == nil {
		return &errorHandler.HttpErr{
			Err:        errors.New("areaPoints is empty"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if len(area.AreaPoints) < 3 {
		return &errorHandler.HttpErr{
			Err:        errors.New("size of areaPoints must be greater than 2"),
			StatusCode: http.StatusBadRequest,
		}
	}
	for _, areaPoint := range area.AreaPoints {
		httpErr := AreaPointValidator.ValidateAreaPoint(&areaPoint)
		if httpErr != nil {
			return httpErr
		}
	}

	return nil
}
