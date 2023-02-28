package repository

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/pkg/paginator"
)

type Animal interface {
	Get(id int) (*entity.Animal, error)
	GetAnimalLocations(animalId int) (*[]entity.Location, error)
	Search(params *filter.AnimalFilterParams) (*[]entity.Animal, error)
	GetAnimalsByAccountId(accountId int) (*[]entity.Animal, error)
}

type AnimalRepository struct {
	Db *gorm.DB
}

func NewAnimalRepository(db *gorm.DB) Animal {
	return &AnimalRepository{Db: db}
}

func (a *AnimalRepository) Get(id int) (*entity.Animal, error) {
	var animal entity.Animal
	err := a.Db.
		Preload("VisitedLocations").
		Preload("AnimalTypes").
		First(&animal, id).Error

	if err != nil {
		return nil, err
	}

	return &animal, nil
}

func (a *AnimalRepository) GetAnimalLocations(animalId int) (*[]entity.Location, error) {
	var animal entity.Animal

	err := a.Db.
		Preload("VisitedLocations").
		Select("Id").
		First(&animal, animalId).Error

	if err != nil {
		return nil, err
	}

	return &animal.VisitedLocations, nil
}

func (a *AnimalRepository) Search(params *filter.AnimalFilterParams) (*[]entity.Animal, error) {
	var animals []entity.Animal
	err := a.Db.
		Scopes(paginator.Paginate(params), filter.AnimalFilter(params)).
		Order("id").
		Find(&animals).Error
	if err != nil {
		return nil, err
	}

	return &animals, nil
}

func (a *AnimalRepository) GetAnimalsByAccountId(accountId int) (*[]entity.Animal, error) {
	var animals []entity.Animal

	err := a.Db.
		Where("chipper_id = ?", accountId).
		Find(&animals).Error
	if err != nil {
		return nil, err
	}

	return &animals, nil
}
