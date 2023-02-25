package repository

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
)

type Animal interface {
	Get(id int) (*entity.Animal, error)
}

type AnimalRepository struct {
	Db *gorm.DB
}

func NewAnimalRepository(db *gorm.DB) Animal {
	return &AnimalRepository{Db: db}
}

func (a *AnimalRepository) Get(id int) (*entity.Animal, error) {
	var animal entity.Animal
	err := a.Db.First(&animal, id).Error
	if err != nil {
		return nil, err
	}

	return &animal, nil
}
