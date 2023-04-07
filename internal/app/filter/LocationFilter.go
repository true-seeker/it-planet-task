package filter

import (
	"errors"
	"it-planet-task/internal/app/validator"
	"it-planet-task/pkg/errorHandler"
	"net/http"
	"net/url"
)

type LocationFilterParams struct {
	Latitude  *float64
	Longitude *float64
}

func NewLocationCoordinatesParams(q url.Values) (*LocationFilterParams, *errorHandler.HttpErr) {
	params := &LocationFilterParams{}
	if q.Get("latitude") == "" {
		return nil, &errorHandler.HttpErr{
			Err:        errors.New("latitude is missing"),
			StatusCode: http.StatusBadRequest,
		}
	}
	if q.Get("longitude") == "" {
		return nil, &errorHandler.HttpErr{
			Err:        errors.New("longitude is missing"),
			StatusCode: http.StatusBadRequest,
		}
	}

	latitude, httpErr := validator.ValidateAndReturnFloatField(q.Get("latitude"), "latitude", 64)
	if httpErr != nil {
		return nil, httpErr
	}
	params.Latitude = &latitude

	longitude, httpErr := validator.ValidateAndReturnFloatField(q.Get("longitude"), "longitude", 64)
	if httpErr != nil {
		return nil, httpErr
	}
	params.Longitude = &longitude
	return params, nil
}
