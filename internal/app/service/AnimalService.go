package service

import (
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
)

type Animal interface {
	Get(id int) (*response.Animal, error)
	GetAnimalLocations(animalId int) (*[]response.Location, error)
	Search(params *filter.AnimalFilterParams) (*[]response.Animal, error)
	GetAnimalsByAccountId(accountId int) (*[]entity.Animal, error)
	GetAnimalsByAnimalTypeId(animalTypeId int) (*[]entity.Animal, error)
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

func (a *AnimalService) Search(params *filter.AnimalFilterParams) (*[]response.Animal, error) {
	var animalResponses *[]response.Animal

	animals, err := a.animalRepo.Search(params)
	if err != nil {
		return nil, err
	}
	animalResponses = mapper.AnimalsToAnimalResponses(animals)

	return animalResponses, nil
}

func (a *AnimalService) GetAnimalsByAccountId(accountId int) (*[]entity.Animal, error) {
	animals, err := a.animalRepo.GetAnimalsByAccountId(accountId)
	if err != nil {
		return nil, err
	}
	return animals, nil
}

func (a *AnimalService) GetAnimalsByAnimalTypeId(animalTypeId int) (*[]entity.Animal, error) {
	animals, err := a.animalRepo.GetAnimalsByAnimalTypeId(animalTypeId)
	if err != nil {
		return nil, err
	}
	return animals, nil
}
