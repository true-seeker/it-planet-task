package repository

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
)

type Area interface {
	Get(id int) (*entity.Area, error)
	Create(area *entity.Area) (*entity.Area, error)
	Update(area *entity.Area) (*entity.Area, error)
	Delete(area *entity.Area) error
}

type AreaRepository struct {
	Db *gorm.DB
}

func NewAreaRepository(db *gorm.DB) Area {
	return &AreaRepository{Db: db}
}

func (a AreaRepository) Get(id int) (*entity.Area, error) {
	//TODO implement me
	panic("implement me")
}

func (a AreaRepository) Create(area *entity.Area) (*entity.Area, error) {
	//TODO implement me
	panic("implement me")
}

func (a AreaRepository) Update(area *entity.Area) (*entity.Area, error) {
	//TODO implement me
	panic("implement me")
}

func (a AreaRepository) Delete(area *entity.Area) error {
	//TODO implement me
	panic("implement me")
}
