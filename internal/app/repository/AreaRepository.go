package repository

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/pkg/paginator"
)

type Area interface {
	Get(id int) (*entity.Area, error)
	Create(area *entity.Area) (*entity.Area, error)
	Update(area *entity.Area) (*entity.Area, error)
	Delete(id int) error
	Search(params *filter.AreaFilterParams) (*[]entity.Area, error)
}

type AreaRepository struct {
	Db *gorm.DB
}

func NewAreaRepository(db *gorm.DB) Area {
	return &AreaRepository{Db: db}
}

func (a *AreaRepository) Get(id int) (*entity.Area, error) {
	var area entity.Area
	err := a.Db.
		Preload("AreaPoints").
		First(&area, id).Error
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
	a.Db.Exec("DELETE FROM area_points WHERE area_id = ?", area.Id)
	err := a.Db.Save(&area).Error
	if err != nil {
		return nil, err
	}

	return area, nil
}

func (a *AreaRepository) Delete(id int) error {
	err := a.Db.Delete(&entity.Area{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *AreaRepository) Search(params *filter.AreaFilterParams) (*[]entity.Area, error) {
	var areas []entity.Area
	err := a.Db.
		Scopes(paginator.Paginate(params)).
		Order("id").
		Find(&areas).Error
	if err != nil {
		return nil, err
	}

	return &areas, nil
}
