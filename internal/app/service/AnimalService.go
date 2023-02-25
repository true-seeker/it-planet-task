package service

import (
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
	"net/url"
)

type Animal interface {
	Get(id int) (*response.Animal, error)
	GetAnimalLocations(animalId int) (*[]response.Location, error)
	Search(query url.Values) (*[]response.Animal, error)
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

	animalResponse = mapper.AnimalToAnimalResponse(animal)

	return animalResponse, nil
}

func (a *AnimalService) GetAnimalLocations(animalId int) (*[]response.Location, error) {
	var locationsResponse *[]response.Location

	locations, err := a.animalRepo.GetAnimalLocations(animalId)
	if err != nil {
		return nil, err
	}

	locationsResponse = mapper.LocationsToLocationResponses(locations)

	return locationsResponse, nil
}

func (a *AnimalService) Search(query url.Values) (*[]response.Animal, error) {
	var animalResponses *[]response.Animal

	animals, err := a.animalRepo.Search(query)
	if err != nil {
		return nil, err
	}
	animalResponses = mapper.AnimalsToAnimalResponses(animals)

	return animalResponses, nil
}
