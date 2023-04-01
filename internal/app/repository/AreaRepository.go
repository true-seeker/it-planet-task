package repository

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
)

type Area interface {
	Get(id int) (*entity.Area, error)
	Create(area *entity.Area) (*entity.Area, error)
	Update(area *entity.Area) (*entity.Area, error)
	Delete(id int) error
}

type AreaRepository struct {
	Db *gorm.DB
}

func NewAreaRepository(db *gorm.DB) Area {
	return &AreaRepository{Db: db}
}

func (a *AreaRepository) Get(id int) (*entity.Area, error) {
	var area entity.Area
	err := a.Db.First(&area, id).Error
	if err != nil {
		return nil, err
	}

	return &area, nil
}

func (a *AreaRepository) Create(area *entity.Area) (*entity.Area, error) {
	err := a.Db.Create(&area).Error
	if err != nil {
		return nil, err
	}

	return area, nil
}

func (a *AreaRepository) Update(area *entity.Area) (*entity.Area, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AreaRepository) Delete(id int) error {
	err := a.Db.Delete(&entity.Area{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
