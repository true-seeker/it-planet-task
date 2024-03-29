package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/input"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
	"it-planet-task/pkg/errorHandler"
	"net/http"
	"time"
)

type Animal interface {
	Get(id int) (*response.Animal, *errorHandler.HttpErr)
	Search(params *filter.AnimalFilterParams) (*[]response.Animal, error)
	GetAnimalsByAccountId(accountId int) (*[]entity.Animal, error)
	GetAnimalsByAnimalTypeId(animalTypeId int) (*[]entity.Animal, error)
	GetAnimalsByLocationId(locationId int) (*[]entity.Animal, error)
	Create(animal *entity.Animal) (*response.Animal, error)
	Update(newAnimal *entity.Animal, oldAnimal *response.Animal) (*response.Animal, error)
	Delete(id int) error
	AddAnimalType(animalId, typeId int) (*response.Animal, error)
	EditAnimalType(animalId int, animalTypeUpdateInput *input.AnimalTypeUpdate) (*response.Animal, error)
	DeleteAnimalType(animalId int, typeId int) (*response.Animal, error)
}

type AnimalService struct {
	animalRepo repository.Animal
}

func NewAnimalService(animalRepo repository.Animal) Animal {
	return &AnimalService{animalRepo: animalRepo}
}

func (a *AnimalService) Get(id int) (*response.Animal, *errorHandler.HttpErr) {
	animalResponse := &response.Animal{}

	animal, err := a.animalRepo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorHandler.NewHttpErr(fmt.Sprintf("Animal with id %d does not exists", id), http.StatusNotFound)
		} else {
			return nil, errorHandler.NewHttpErr(err.Error(), http.StatusBadRequest)
		}
	}

	animalResponse = mapper.AnimalToAnimalResponse(animal)

	return animalResponse, nil
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

	animal.LifeStatus = entity.Alive
	animal.ChippingDateTime = time.Now()
	animal.DeathDateTime = nil

	animal, err := a.animalRepo.Create(animal)
	if err != nil {
		return nil, err
	}

	animalResponse = mapper.AnimalToAnimalResponse(animal)

	return animalResponse, nil
}

func (a *AnimalService) Update(newAnimal *entity.Animal, oldAnimal *response.Animal) (*response.Animal, error) {
	animalResponse := &response.Animal{}

	if oldAnimal.LifeStatus == entity.Alive && newAnimal.LifeStatus == entity.Dead {
		deathDateTime := time.Now()
		newAnimal.DeathDateTime = &deathDateTime
	} else {
		newAnimal.DeathDateTime = oldAnimal.DeathDateTime
	}
	newAnimal.ChippingDateTime = oldAnimal.ChippingDateTime

	newAnimal, err := a.animalRepo.Update(newAnimal)
	if err != nil {
		return nil, err
	}

	animalResponse = mapper.AnimalToAnimalResponse(newAnimal)

	return animalResponse, nil
}

func (a *AnimalService) Delete(id int) error {
	return a.animalRepo.Delete(id)
}

func (a *AnimalService) AddAnimalType(animalId, typeId int) (*response.Animal, error) {
	animalResponse := &response.Animal{}
	animal, err := a.animalRepo.AddAnimalType(animalId, typeId)
	if err != nil {
		return nil, err
	}

	animalResponse = mapper.AnimalToAnimalResponse(animal)

	return animalResponse, nil
}

func (a *AnimalService) EditAnimalType(animalId int, animalTypeUpdateInput *input.AnimalTypeUpdate) (*response.Animal, error) {
	animalResponse := &response.Animal{}
	animal, err := a.animalRepo.EditAnimalType(animalId, animalTypeUpdateInput)
	if err != nil {
		return nil, err
	}

	animalResponse = mapper.AnimalToAnimalResponse(animal)

	return animalResponse, nil
}

func (a *AnimalService) DeleteAnimalType(animalId int, typeId int) (*response.Animal, error) {
	animalResponse := &response.Animal{}
	animal, err := a.animalRepo.DeleteAnimalType(animalId, typeId)
	if err != nil {
		return nil, err
	}

	animalResponse = mapper.AnimalToAnimalResponse(animal)

	return animalResponse, nil
}
