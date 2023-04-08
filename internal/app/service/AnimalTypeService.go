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

type AnimalType interface {
	Get(id int) (*response.AnimalType, *errorHandler.HttpErr)
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

func (a *AnimalTypeService) Get(id int) (*response.AnimalType, *errorHandler.HttpErr) {
	animalTypeResponse := &response.AnimalType{}

	animalType, err := a.animalTypeRepo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorHandler.NewHttpErr(fmt.Sprintf("Account type with id %d does not exists", id), http.StatusNotFound)
		} else {
			return nil, errorHandler.NewHttpErr(err.Error(), http.StatusBadRequest)
		}
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
