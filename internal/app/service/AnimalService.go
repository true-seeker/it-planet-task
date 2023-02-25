package service

import (
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
)

type Animal interface {
	Get(id int) (*response.Animal, error)
	GetAnimalLocations(animalId int) (*[]response.Location, error)
}

type AnimalService struct {
	animalRepo repository.Animal
}

func NewAnimalService(animalRepo repository.Animal) Animal {
	return &AnimalService{animalRepo: animalRepo}
}

func (a *AnimalService) Get(id int) (*response.Animal, error) {
	animalResponse := &response.Animal{}

	animal, err := a.animalRepo.Get(id)
	if err != nil {
		return nil, err
	}

	animalResponse.Id = animal.Id
	animalResponse.Weight = animal.Weight
	animalResponse.Length = animal.Length
	animalResponse.Gender = animal.Gender
	animalResponse.Height = animal.Height
	animalResponse.LifeStatus = animal.LifeStatus
	animalResponse.ChippingDateTime = animal.ChippingDateTime
	animalResponse.ChipperId = animal.ChipperId
	animalResponse.ChippingLocationId = animal.ChippingLocationId
	animalResponse.DeathDateTime = animal.DeathDateTime
	// TODO mapper
	return animalResponse, nil
}

func (a *AnimalService) GetAnimalLocations(animalId int) (*[]response.Location, error) {
	var locationsResponse []response.Location

	locations, err := a.animalRepo.GetAnimalLocations(animalId)
	if err != nil {
		return nil, err
	}
	for _, location := range *locations {
		locationResponse := response.Location{}
		locationResponse.Id = location.Id
		locationResponse.Latitude = location.Latitude
		locationResponse.Longitude = location.Longitude
		locationsResponse = append(locationsResponse, locationResponse)
	}

	return &locationsResponse, nil
}
