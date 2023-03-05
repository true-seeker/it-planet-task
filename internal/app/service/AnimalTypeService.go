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
	GetByType(animalType *entity.AnimalType) *entity.AnimalType
	GetByIds(ids *[]int) (*[]response.AnimalType, error)
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
	animalTypeResponse := &response.AnimalType{}

	animalType, err := a.animalTypeRepo.Create(animalType)
	if err != nil {
		return nil, err
	}

	animalTypeResponse = mapper.AnimalTypeToAnimalTypeResponse(animalType)

	return animalTypeResponse, nil
}

func (a *AnimalTypeService) Update(animalType *entity.AnimalType) (*response.AnimalType, error) {
	animalTypeResponse := &response.AnimalType{}

	animalType, err := a.animalTypeRepo.Update(animalType)
	if err != nil {
		return nil, err
	}

	animalTypeResponse = mapper.AnimalTypeToAnimalTypeResponse(animalType)

	return animalTypeResponse, nil
}

func (a *AnimalTypeService) Delete(animalTypeId int) error {
	return a.animalTypeRepo.Delete(animalTypeId)
}

func (a *AnimalTypeService) GetByType(animalType *entity.AnimalType) *entity.AnimalType {
	return a.animalTypeRepo.GetByType(animalType)
}

func (a *AnimalTypeService) GetByIds(ids *[]int) (*[]response.AnimalType, error) {
	animalTypeResponses := &[]response.AnimalType{}

	animalTypes, err := a.animalTypeRepo.GetByIds(ids)
	if err != nil {
		return nil, err
	}

	animalTypeResponses = mapper.AnimalTypesToAnimalTypeResponses(animalTypes)

	return animalTypeResponses, nil
}
