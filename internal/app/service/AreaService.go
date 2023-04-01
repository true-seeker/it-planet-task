package service

import (
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
)

type Area interface {
	Get(id int) (*entity.Area, error)
	Create(area *entity.Area) (*response.Area, error)
	Update(area *entity.Area) (*response.Area, error)
	Delete(area *entity.Area) error
}

type AreaService struct {
	areaRepo repository.Area
}

func NewAreaService(areaRepo repository.Area) Area {
	return &AreaService{areaRepo: areaRepo}
}

func (a AreaService) Get(id int) (*entity.Area, error) {
	//TODO implement me
	panic("implement me")
}

func (a AreaService) Create(area *entity.Area) (*response.Area, error) {
	//TODO implement me
	panic("implement me")
}

func (a AreaService) Update(area *entity.Area) (*response.Area, error) {
	//TODO implement me
	panic("implement me")
}

func (a AreaService) Delete(area *entity.Area) error {
	//TODO implement me
	panic("implement me")
}
