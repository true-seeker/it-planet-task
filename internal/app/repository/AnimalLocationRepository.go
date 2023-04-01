package repository

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/pkg/paginator"
)

type AnimalLocation interface {
	GetAnimalLocations(animalId int, params *filter.AnimalLocationFilterParams) (*[]entity.AnimalLocation, error)
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

func (a *AnimalLocationRepository) GetAnimalLocations(animalId int, params *filter.AnimalLocationFilterParams) (*[]entity.AnimalLocation, error) {
	var animalLocations []entity.AnimalLocation

	err := a.Db.Where("animal_id = ?", animalId).
		Scopes(paginator.Paginate(params),
			filter.AnimalLocationFilter(params)).
		Find(&animalLocations).Error

	if err != nil {
		return nil, err
	}

	return &animalLocations, nil
}

func (a *AnimalLocationRepository) AddAnimalLocationPoint(newAnimalLocation *entity.AnimalLocation) (*entity.AnimalLocation, error) {
	a.Db.Save(newAnimalLocation)

	return a.Get(newAnimalLocation.Id)
}

func (a *AnimalLocationRepository) EditAnimalLocationPoint(visitedLocationPointId int, locationPointId int) (*entity.AnimalLocation, error) {
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
