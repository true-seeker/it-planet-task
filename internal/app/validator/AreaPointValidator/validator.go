package AreaPointValidator

import (
	"errors"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

func ValidateAreaPoint(areaPoint *entity.AreaPoint) *errorHandler.HttpErr {
	if areaPoint.Latitude == nil || *areaPoint.Latitude < -90 || *areaPoint.Latitude > 90 {
		return &errorHandler.HttpErr{
			Err:        errors.New("invalid latitude"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if areaPoint.Longitude == nil || *areaPoint.Longitude < -180 || *areaPoint.Longitude > 180 {
		return &errorHandler.HttpErr{
			Err:        errors.New("invalid longitude"),
			StatusCode: http.StatusBadRequest,
		}
	}
	return nil
}
