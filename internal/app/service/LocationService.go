package service

import (
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
)

type Location interface {
	Get(id int) (*response.Location, error)
}

type LocationService struct {
	locationRepo repository.Location
}

func NewLocationService(serviceRepo repository.Location) Location {
	return &LocationService{locationRepo: serviceRepo}
}

func (l *LocationService) Get(id int) (*response.Location, error) {
	locationResponse := &response.Location{}

	location, err := l.locationRepo.Get(id)
	if err != nil {
		return nil, err
	}

	locationResponse.Id = location.Id
	locationResponse.Latitude = location.Latitude
	locationResponse.Longitude = location.Longitude

	return locationResponse, nil
}
