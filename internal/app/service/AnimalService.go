package service

import (
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
	"it-planet-task/internal/app/validator/AnimalValidator"
	"time"
)

type Animal interface {
	Get(id int) (*response.Animal, error)
	GetAnimalLocations(animalId int) (*[]response.Location, error)
	Search(params *filter.AnimalFilterParams) (*[]response.Animal, error)
	GetAnimalsByAccountId(accountId int) (*[]entity.Animal, error)
	GetAnimalsByAnimalTypeId(animalTypeId int) (*[]entity.Animal, error)
	GetAnimalsByLocationId(locationId int) (*[]entity.Animal, error)
	Create(animal *entity.Animal) (*response.Animal, error)
	Update(animal *entity.Animal) (*response.Animal, error)
	Delete(id int) error
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

func (a *AnimalService) GetAnimalsByLocationId(locationId int) (*[]entity.Animal, error) {
	animals, err := a.animalRepo.GetAnimalsByLocationId(locationId)
	if err != nil {
		return nil, err
	}
	return animals, nil
}

func (a *AnimalService) Create(animal *entity.Animal) (*response.Animal, error) {
	animalResponse := &response.Animal{}

	animal.LifeStatus = AnimalValidator.Alive
	animal.ChippingDateTime = time.Now()
	animal.DeathDateTime = nil

	animal, err := a.animalRepo.Create(animal)
	if err != nil {
		return nil, err
	}

	animalResponse = mapper.AnimalToAnimalResponse(animal)

	return animalResponse, nil
}

func (a *AnimalService) Update(animal *entity.Animal) (*response.Animal, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AnimalService) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
