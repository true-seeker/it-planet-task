package repository

import (
	"fmt"
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/input"
	"it-planet-task/pkg/paginator"
)

type Animal interface {
	Get(id int) (*entity.Animal, error)
	GetAnimalLocations(animalId int) (*[]entity.AnimalLocation, error)
	Search(params *filter.AnimalFilterParams) (*[]entity.Animal, error)
	GetAnimalsByAccountId(accountId int) (*[]entity.Animal, error)
	GetAnimalsByAnimalTypeId(accountId int) (*[]entity.Animal, error)
	GetAnimalsByLocationId(locationId int) (*[]entity.Animal, error)
	Create(animal *entity.Animal) (*entity.Animal, error)
	Update(animal *entity.Animal) (*entity.Animal, error)
	Delete(id int) error
	AddAnimalType(animalId, typeId int) (*entity.Animal, error)
	EditAnimalType(animalId int, input *input.AnimalTypeUpdate) (*entity.Animal, error)
	DeleteAnimalType(animalId int, typeId int) (*entity.Animal, error)
	AddAnimalLocationPoint(animalId int, newAnimalLocation *entity.AnimalLocation) (*entity.AnimalLocation, error)
	EditAnimalLocationPoint(visitedLocationPointId int, locationPointId int) (*entity.AnimalLocation, error)
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

func (a *AnimalRepository) GetAnimalLocations(animalId int) (*[]entity.AnimalLocation, error) {
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

func (a *AnimalRepository) GetAnimalsByAnimalTypeId(animalTypeId int) (*[]entity.Animal, error) {
	var animals []entity.Animal
	err := a.Db.Joins("left outer join animal_animal_type on animal_id=animals.id"+
		" left outer join animal_types on animal_type_id=animal_types.id ").
		Where("animal_type_id = ?", animalTypeId).
		Find(&animals).Error
	if err != nil {
		return nil, err
	}

	return &animals, nil
}

func (a *AnimalRepository) GetAnimalsByLocationId(locationId int) (*[]entity.Animal, error) {
	var animals []entity.Animal
	err := a.Db.Joins("left outer join animal_visited_locations on animal_id=animals.id").
		Where("location_id = ?", locationId).
		Find(&animals).Error
	if err != nil {
		return nil, err
	}

	return &animals, nil
}

func (a *AnimalRepository) Create(animal *entity.Animal) (*entity.Animal, error) {
	err := a.Db.Create(&animal).Error
	if err != nil {
		return nil, err
	}

	return animal, nil
}

func (a *AnimalRepository) Update(animal *entity.Animal) (*entity.Animal, error) {
	err := a.Db.Save(&animal).Error
	if err != nil {
		return nil, err
	}

	return animal, nil
}

func (a *AnimalRepository) Delete(id int) error {
	err := a.Db.Delete(&entity.Animal{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *AnimalRepository) AddAnimalType(animalId, typeId int) (*entity.Animal, error) {
	a.Db.Exec("INSERT INTO animal_animal_type(animal_id, animal_type_id) VALUES (?,?)", animalId, typeId)
	return a.Get(animalId)
}

func (a *AnimalRepository) EditAnimalType(animalId int, input *input.AnimalTypeUpdate) (*entity.Animal, error) {
	a.Db.Exec("UPDATE animal_animal_type SET animal_type_id = ? WHERE animal_id = ? AND animal_type_id = ?", input.NewTypeId, animalId, input.OldTypeId)
	return a.Get(animalId)
}

func (a *AnimalRepository) DeleteAnimalType(animalId int, typeId int) (*entity.Animal, error) {
	a.Db.Exec("DELETE FROM animal_animal_type WHERE animal_id = ? AND animal_type_id = ?", animalId, typeId)
	return a.Get(animalId)
}

func (a *AnimalRepository) AddAnimalLocationPoint(animalId int, newAnimalLocation *entity.AnimalLocation) (*entity.AnimalLocation, error) {
	a.Db.Save(newAnimalLocation)

	a.Db.Exec("INSERT INTO animal_visited_locations(animal_id, animal_location_id) VALUES (?,?)", animalId, newAnimalLocation.Id)

	return newAnimalLocation, nil
}

func (a *AnimalRepository) EditAnimalLocationPoint(visitedLocationPointId int, locationPointId int) (*entity.AnimalLocation, error) {
	fmt.Println(locationPointId, visitedLocationPointId)
	a.Db.Exec("UPDATE animal_locations SET location_point_id = ? WHERE id = ?", locationPointId, visitedLocationPointId)
	// TODO разбить на отдельный сервис и репозиторий
	var animalLocation entity.AnimalLocation
	a.Db.First(&animalLocation, visitedLocationPointId)
	return &animalLocation, nil
}
