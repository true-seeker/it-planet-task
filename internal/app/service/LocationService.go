package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

type Location interface {
	Get(id int) (*response.Location, *errorHandler.HttpErr)
	Create(location *entity.Location) (*response.Location, error)
	Update(location *entity.Location) (*response.Location, error)
	Delete(id int) error
	GetByCoordinates(location *entity.Location) (*response.Location, *errorHandler.HttpErr)
}

type LocationService struct {
	locationRepo repository.Location
}

func NewLocationService(serviceRepo repository.Location) Location {
	return &LocationService{locationRepo: serviceRepo}
}

func (l *LocationService) Get(id int) (*response.Location, *errorHandler.HttpErr) {
	locationResponse := &response.Location{}

	location, err := l.locationRepo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorHandler.NewHttpErr(fmt.Sprintf("Location with id %d does not exists", id), http.StatusNotFound)
		} else {
			return nil, errorHandler.NewHttpErr(err.Error(), http.StatusBadRequest)
		}
	}

	locationResponse = mapper.LocationToLocationResponse(location)

	return locationResponse, nil
}

func (l *LocationService) Create(location *entity.Location) (*response.Location, error) {
	locationResponse := &response.Location{}

	location, err := l.locationRepo.Create(location)
	if err != nil {
		return nil, err
	}

	locationResponse = mapper.LocationToLocationResponse(location)

	return locationResponse, nil
}

func (l *LocationService) Update(location *entity.Location) (*response.Location, error) {
	locationResponse := &response.Location{}

	location, err := l.locationRepo.Update(location)
	if err != nil {
		return nil, err
	}

	locationResponse = mapper.LocationToLocationResponse(location)

	return locationResponse, nil
}

func (l *LocationService) Delete(id int) error {
	return l.locationRepo.Delete(id)
}

func (l *LocationService) GetByCoordinates(location *entity.Location) (*response.Location, *errorHandler.HttpErr) {
	locationResponse := &response.Location{}
	location, err := l.locationRepo.GetByCoordinates(location)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorHandler.NewHttpErr("location with these coordinates does not exists", http.StatusNotFound)
		} else {
			return nil, errorHandler.NewHttpErr(err.Error(), http.StatusBadRequest)
		}
	}

	locationResponse = mapper.LocationToLocationResponse(location)

	return locationResponse, nil
}
