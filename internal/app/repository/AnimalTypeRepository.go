package repository

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
)

type AnimalType interface {
	Get(id int) (*entity.AnimalType, error)
	Create(animalType *entity.AnimalType) (*entity.AnimalType, error)
	Update(animalType *entity.AnimalType) (*entity.AnimalType, error)
	Delete(animalTypeId int) error
}

type AnimalTypeRepository struct {
	Db *gorm.DB
}

func NewAnimalTypeRepository(db *gorm.DB) AnimalType {
	return &AnimalTypeRepository{Db: db}
}

func (a *AnimalTypeRepository) Get(id int) (*entity.AnimalType, error) {
	var animalType entity.AnimalType
	err := a.Db.First(&animalType, id).Error
	if err != nil {
		return nil, err
	}

	return &animalType, nil
}

func (a *AnimalTypeRepository) Create(animalType *entity.AnimalType) (*entity.AnimalType, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AnimalTypeRepository) Update(animalType *entity.AnimalType) (*entity.AnimalType, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AnimalTypeRepository) Delete(animalTypeId int) error {
	//TODO implement me
	panic("implement me")
}
