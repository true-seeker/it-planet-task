package AreaValidator

import (
	"errors"
	"fmt"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service"
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

	polygonCandidate := service.NewPolygon(&area.AreaPoints)

	if !polygonCandidate.IsConvex() {
		return &errorHandler.HttpErr{
			Err:        errors.New("area must be convex"),
			StatusCode: http.StatusBadRequest,
		}
	}

	return nil
}

func ValidateIntersection(area *entity.Area, existingArea *entity.Area) *errorHandler.HttpErr {
	polygonCandidate := service.NewPolygon(&area.AreaPoints)
	existingPolygon := service.NewPolygon(&existingArea.AreaPoints)
	if polygonCandidate.IsIntersect(existingPolygon) {
		return &errorHandler.HttpErr{
			Err:        errors.New(fmt.Sprintf("area intersect with existing area with id %d", existingArea.Id)),
			StatusCode: http.StatusBadRequest,
		}
	}

	return nil
}
