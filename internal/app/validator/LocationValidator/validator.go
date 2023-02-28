package LocationValidator

import (
	"errors"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

func ValidateLocation(location *entity.Location) *errorHandler.HttpErr {
	if location.Latitude == nil || *location.Latitude < -90 || *location.Latitude > 90 {
		return &errorHandler.HttpErr{
			Err:        errors.New("invalid latitude"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if location.Longitude == nil || *location.Longitude < -180 || *location.Longitude > 180 {
		return &errorHandler.HttpErr{
			Err:        errors.New("invalid longitude"),
			StatusCode: http.StatusBadRequest,
		}
	}
	return nil
}
