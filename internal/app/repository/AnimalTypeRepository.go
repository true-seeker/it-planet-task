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
	GetByType(animalType *entity.AnimalType) *entity.AnimalType
	GetByIds(ids *[]int) (*[]entity.AnimalType, error)
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
	err := a.Db.Create(&animalType).Error
	if err != nil {
		return nil, err
	}

	return animalType, nil
}

func (a *AnimalTypeRepository) Update(animalType *entity.AnimalType) (*entity.AnimalType, error) {
	err := a.Db.Save(&animalType).Error
	if err != nil {
		return nil, err
	}

	return animalType, nil
}

func (a *AnimalTypeRepository) Delete(animalTypeId int) error {
	err := a.Db.Delete(&entity.AnimalType{}, animalTypeId).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *AnimalTypeRepository) GetByType(animalType *entity.AnimalType) *entity.AnimalType {
	ant := &entity.AnimalType{}
	a.Db.Where("type = ?", animalType.Type).First(ant)
	return ant
}

func (a *AnimalTypeRepository) GetByIds(ids *[]int) (*[]entity.AnimalType, error) {
	ants := &[]entity.AnimalType{}
	err := a.Db.Find(ants, ids).Error
	if err != nil {
		return nil, err
	}
	return ants, nil
}
