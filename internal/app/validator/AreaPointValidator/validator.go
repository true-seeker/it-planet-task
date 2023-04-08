package AreaPointValidator

import (
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

func ValidateAreaPoint(areaPoint *entity.AreaPoint) *errorHandler.HttpErr {
	if areaPoint.Latitude == nil || *areaPoint.Latitude < -90 || *areaPoint.Latitude > 90 {
		return errorHandler.NewHttpErr("invalid latitude", http.StatusBadRequest)
	}

	if areaPoint.Longitude == nil || *areaPoint.Longitude < -180 || *areaPoint.Longitude > 180 {
		return errorHandler.NewHttpErr("invalid longitude", http.StatusBadRequest)
	}
	return nil
}
