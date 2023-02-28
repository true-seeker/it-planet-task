package repository

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
)

type Location interface {
	Get(id int) (*entity.Location, error)
	Create(location *entity.Location) (*entity.Location, error)
	Update(location *entity.Location) (*entity.Location, error)
	Delete(id int) error
	GetByCords(location *entity.Location) (*entity.Location, error)
}

type LocationRepository struct {
	Db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) Location {
	return &LocationRepository{Db: db}
}

func (a *LocationRepository) Get(id int) (*entity.Location, error) {
	var location entity.Location
	err := a.Db.First(&location, id).Error
	if err != nil {
		return nil, err
	}

	return &location, nil
}

func (a *LocationRepository) Create(location *entity.Location) (*entity.Location, error) {
	err := a.Db.Create(&location).Error
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (a *LocationRepository) Update(location *entity.Location) (*entity.Location, error) {
	err := a.Db.Save(&location).Error
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (a *LocationRepository) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (a *LocationRepository) GetByCords(location *entity.Location) (*entity.Location, error) {
	lc := &entity.Location{}
	a.Db.Where("longitude = ? AND latitude = ?", location.Longitude, location.Latitude).First(lc)
	return lc, nil
}
