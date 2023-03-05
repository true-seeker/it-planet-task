package service

import (
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
	"time"
)

type AnimalLocation interface {
	GetAnimalLocations(animalId int) (*[]response.AnimalLocation, error)
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

func (a *AnimalLocationService) GetAnimalLocations(animalId int) (*[]response.AnimalLocation, error) {
	var animalLocationsResponse *[]response.AnimalLocation

	animalLocations, err := a.animalLocationRepo.GetAnimalLocations(animalId)
	if err != nil {
		return nil, err
	}

	animalLocationsResponse = mapper.AnimalLocationsToAnimalLocationResponses(animalLocations)

	return animalLocationsResponse, nil
}

func (a *AnimalLocationService) AddAnimalLocationPoint(animalId int, pointId int) (*response.AnimalLocation, error) {
	animalLocationResponse := &response.AnimalLocation{}

	animalLocation := &entity.AnimalLocation{
		DateTimeOfVisitLocationPoint: time.Now(),
		LocationPointId:              pointId,
	}

	animalLocation, err := a.animalLocationRepo.AddAnimalLocationPoint(animalId, animalLocation)
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
