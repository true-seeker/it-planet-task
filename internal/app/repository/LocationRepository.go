package repository

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
)

type Location interface {
	Get(id int) (*entity.Location, error)
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
