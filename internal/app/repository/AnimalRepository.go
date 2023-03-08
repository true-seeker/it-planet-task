package repository

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/input"
	"it-planet-task/pkg/paginator"
)

type Animal interface {
	Get(id int) (*entity.Animal, error)
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

func (a *AnimalRepository) Search(params *filter.AnimalFilterParams) (*[]entity.Animal, error) {
	var animals []entity.Animal
	err := a.Db.
		Scopes(paginator.Paginate(params), filter.AnimalFilter(params)).
		Order("id").
		Preload("AnimalTypes").
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
	err := a.Db.Joins("join animal_animal_type on animal_id=animals.id"+
		" join animal_types on animal_type_id=animal_types.id ").
		Where("animal_type_id = ?", animalTypeId).
		Find(&animals).Error
	if err != nil {
		return nil, err
	}

	return &animals, nil
}

func (a *AnimalRepository) GetAnimalsByLocationId(locationId int) (*[]entity.Animal, error) {
	var animals []entity.Animal
	err := a.Db.Joins("join animal_visited_locations on animal_id=animals.id "+
		"join animal_locations on animal_location_id=animal_locations.id ").
		Where("location_point_id = ?", locationId).
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

	return a.Get(animal.Id)
}

func (a *AnimalRepository) Update(animal *entity.Animal) (*entity.Animal, error) {
	err := a.Db.Save(&animal).Error
	if err != nil {
		return nil, err
	}

	return a.Get(animal.Id)
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
