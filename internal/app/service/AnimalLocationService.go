package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
	"it-planet-task/pkg/errorHandler"
	"net/http"
	"time"
)

type AnimalLocation interface {
	Get(id int) (*response.AnimalLocation, *errorHandler.HttpErr)
	GetAnimalLocations(animalId int, params *filter.AnimalLocationFilterParams) (*[]response.AnimalLocation, *errorHandler.HttpErr)
	AddAnimalLocationPoint(animalId int, pointId int) (*response.AnimalLocation, error)
	EditAnimalLocationPoint(visitedLocationPointId int, locationPointId int) (*response.AnimalLocation, error)
	DeleteAnimalLocationPoint(visitedPointId int) error
}

type AnimalLocationService struct {
	animalLocationRepo repository.AnimalLocation
}

func NewAnimalLocationService(animalLocationRepo repository.AnimalLocation) AnimalLocation {
	return &AnimalLocationService{animalLocationRepo: animalLocationRepo}
}

func (a *AnimalLocationService) GetAnimalLocations(animalId int, params *filter.AnimalLocationFilterParams) (*[]response.AnimalLocation, *errorHandler.HttpErr) {
	var animalLocationsResponse *[]response.AnimalLocation

	animalLocations, err := a.animalLocationRepo.GetAnimalLocations(animalId, params)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorHandler.HttpErr{
				Err:        errors.New(fmt.Sprintf("Animal with id %d does not exists", animalId)),
				StatusCode: http.StatusNotFound,
			}
		} else {
			return nil, &errorHandler.HttpErr{
				Err:        err,
				StatusCode: http.StatusBadRequest,
			}
		}
	}

	animalLocationsResponse = mapper.AnimalLocationsToAnimalLocationResponses(animalLocations)

	return animalLocationsResponse, nil
}

func (a *AnimalLocationService) AddAnimalLocationPoint(animalId int, pointId int) (*response.AnimalLocation, error) {
	animalLocationResponse := &response.AnimalLocation{}

	animalLocation := &entity.AnimalLocation{
		DateTimeOfVisitLocationPoint: time.Now(),
		LocationPointId:              pointId,
		AnimalId:                     animalId,
	}

	animalLocation, err := a.animalLocationRepo.AddAnimalLocationPoint(animalLocation)
	if err != nil {
		return nil, err
	}

	animalLocationResponse = mapper.AnimalLocationToAnimalLocationResponse(animalLocation)

	return animalLocationResponse, nil
}

func (a *AnimalLocationService) EditAnimalLocationPoint(visitedLocationPointId int, locationPointId int) (*response.AnimalLocation, error) {
	animalLocationResponse := &response.AnimalLocation{}

	animalLocation, err := a.animalLocationRepo.EditAnimalLocationPoint(visitedLocationPointId, locationPointId)
	if err != nil {
		return nil, err
	}

	animalLocationResponse = mapper.AnimalLocationToAnimalLocationResponse(animalLocation)

	return animalLocationResponse, nil
}

func (a *AnimalLocationService) DeleteAnimalLocationPoint(visitedPointId int) error {
	return a.animalLocationRepo.DeleteAnimalLocationPoint(visitedPointId)
}

func (a *AnimalLocationService) Get(id int) (*response.AnimalLocation, *errorHandler.HttpErr) {
	animalLocationResponse := &response.AnimalLocation{}

	animalLocation, err := a.animalLocationRepo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorHandler.HttpErr{
				Err:        errors.New(fmt.Sprintf("Animal with id %d does not exists", id)),
				StatusCode: http.StatusNotFound,
			}
		} else {
			return nil, &errorHandler.HttpErr{
				Err:        err,
				StatusCode: http.StatusBadRequest,
			}
		}
	}

	animalLocationResponse = mapper.AnimalLocationToAnimalLocationResponse(animalLocation)

	return animalLocationResponse, nil
}
