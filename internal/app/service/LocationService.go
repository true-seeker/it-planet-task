package service

import (
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
)

type Location interface {
	Get(id int) (*response.Location, error)
	Create(location *entity.Location) (*response.Location, error)
	Update(location *entity.Location) (*response.Location, error)
	Delete(id int) error
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

	locationResponse = mapper.LocationToLocationResponse(location)

	return locationResponse, nil
}

func (l *LocationService) Create(location *entity.Location) (*response.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LocationService) Update(location *entity.Location) (*response.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LocationService) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
