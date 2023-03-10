package repository

import (
	"fmt"
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
)

type AnimalLocation interface {
	GetAnimalLocations(animalId int) (*[]entity.AnimalLocation, error)
	AddAnimalLocationPoint(newAnimalLocation *entity.AnimalLocation) (*entity.AnimalLocation, error)
	EditAnimalLocationPoint(visitedLocationPointId int, locationPointId int) (*entity.AnimalLocation, error)
	DeleteAnimalLocationPoint(id int) error
	Get(id int) (*entity.AnimalLocation, error)
}

type AnimalLocationRepository struct {
	Db *gorm.DB
}

func NewAnimalLocationRepository(db *gorm.DB) AnimalLocation {
	return &AnimalLocationRepository{Db: db}
}

func (a *AnimalLocationRepository) GetAnimalLocations(animalId int) (*[]entity.AnimalLocation, error) {
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

func (a *AnimalLocationRepository) AddAnimalLocationPoint(newAnimalLocation *entity.AnimalLocation) (*entity.AnimalLocation, error) {
	a.Db.Save(newAnimalLocation)

	return a.Get(newAnimalLocation.Id)
}

func (a *AnimalLocationRepository) EditAnimalLocationPoint(visitedLocationPointId int, locationPointId int) (*entity.AnimalLocation, error) {
	fmt.Println(locationPointId, visitedLocationPointId)
	a.Db.Exec("UPDATE animal_locations SET location_point_id = ? WHERE id = ?", locationPointId, visitedLocationPointId)

	return a.Get(visitedLocationPointId)
}

func (a *AnimalLocationRepository) DeleteAnimalLocationPoint(id int) error {
	err := a.Db.Delete(&entity.AnimalLocation{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *AnimalLocationRepository) Get(id int) (*entity.AnimalLocation, error) {
	var animalLocation entity.AnimalLocation
	err := a.Db.First(&animalLocation, id).Error

	if err != nil {
		return nil, err
	}

	return &animalLocation, nil
}
