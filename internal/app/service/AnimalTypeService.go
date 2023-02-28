package service

import (
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
)

type AnimalType interface {
	Get(id int) (*response.AnimalType, error)
	Create(animalType *entity.AnimalType) (*response.AnimalType, error)
	Update(animalType *entity.AnimalType) (*response.AnimalType, error)
	Delete(animalTypeId int) error
}

type AnimalTypeService struct {
	animalTypeRepo repository.AnimalType
}

func NewAnimalTypeService(animalTypeRepo repository.AnimalType) AnimalType {
	return &AnimalTypeService{animalTypeRepo: animalTypeRepo}
}

func (a *AnimalTypeService) Get(id int) (*response.AnimalType, error) {
	animalTypeResponse := &response.AnimalType{}

	animalType, err := a.animalTypeRepo.Get(id)
	if err != nil {
		return nil, err
	}

	animalTypeResponse = mapper.AnimalTypeToAnimalTypeResponse(animalType)

	return animalTypeResponse, nil
}

func (a *AnimalTypeService) Create(animalType *entity.AnimalType) (*response.AnimalType, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AnimalTypeService) Update(animalType *entity.AnimalType) (*response.AnimalType, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AnimalTypeService) Delete(animalTypeId int) error {
	//TODO implement me
	panic("implement me")
}
