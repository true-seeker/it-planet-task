package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

type Area interface {
	Get(id int) (*response.Area, *errorHandler.HttpErr)
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

func (a *AreaService) Get(id int) (*response.Area, *errorHandler.HttpErr) {
	areaResponse := &response.Area{}

	area, err := a.areaRepo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorHandler.HttpErr{
				Err:        errors.New(fmt.Sprintf("Area with id %d does not exists", id)),
				StatusCode: http.StatusNotFound,
			}
		} else {
			return nil, &errorHandler.HttpErr{
				Err:        err,
				StatusCode: http.StatusBadRequest,
			}
		}
	}

	areaResponse = mapper.AreaToAreaResponse(area)

	return areaResponse, nil
}

func (a *AreaService) Create(area *entity.Area) (*response.Area, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AreaService) Update(area *entity.Area) (*response.Area, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AreaService) Delete(area *entity.Area) error {
	//TODO implement me
	panic("implement me")
}
