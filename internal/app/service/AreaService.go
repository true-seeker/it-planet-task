package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
	"it-planet-task/internal/app/validator/AreaValidator"
	"it-planet-task/pkg/errorHandler"
	"net/http"
	"net/url"
)

type Area interface {
	Get(id int) (*response.Area, *errorHandler.HttpErr)
	Create(area *entity.Area) (*response.Area, *errorHandler.HttpErr)
	Update(area *entity.Area) (*response.Area, *errorHandler.HttpErr)
	Delete(id int) error
	Search(params *filter.AreaFilterParams) (*[]response.Area, *errorHandler.HttpErr)
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

func (a *AreaService) Create(area *entity.Area) (*response.Area, *errorHandler.HttpErr) {
	from := 0
	size := 10
	for {
		query := fmt.Sprintf("size=%d&from=%d", size, from)
		values, _ := url.ParseQuery(query)
		params, _ := filter.NewAreaFilterParams(values)
		existingAreas, err := a.areaRepo.Search(params)

		if err != nil {
			return nil, &errorHandler.HttpErr{
				Err:        err,
				StatusCode: http.StatusBadRequest,
			}
		}
		if len(*existingAreas) == 0 {
			break
		}
		for _, existingArea := range *existingAreas {
			httpErr := AreaValidator.ValidateIntersectionAndAreaRepeats(area, &existingArea)
			if httpErr != nil {
				return nil, httpErr
			}
		}
		from += size
	}

	areaResponse := &response.Area{}

	area, err := a.areaRepo.Create(area)
	if err != nil {
		return nil, &errorHandler.HttpErr{
			Err:        err,
			StatusCode: http.StatusBadRequest,
		}
	}

	areaResponse = mapper.AreaToAreaResponse(area)

	return areaResponse, nil
}

func (a *AreaService) Update(area *entity.Area) (*response.Area, *errorHandler.HttpErr) {
	from := 0
	size := 10
	for {
		query := fmt.Sprintf("size=%d&from=%d", size, from)
		values, _ := url.ParseQuery(query)
		params, _ := filter.NewAreaFilterParams(values)
		existingAreas, err := a.areaRepo.Search(params)

		if err != nil {
			return nil, &errorHandler.HttpErr{
				Err:        err,
				StatusCode: http.StatusBadRequest,
			}
		}
		if len(*existingAreas) == 0 {
			break
		}
		for _, existingArea := range *existingAreas {
			if existingArea.Id == area.Id {
				continue
			}

			httpErr := AreaValidator.ValidateIntersectionAndAreaRepeats(area, &existingArea)
			if httpErr != nil {
				return nil, httpErr
			}
		}
		from += size
	}

	areaResponse := &response.Area{}

	area, err := a.areaRepo.Update(area)
	if err != nil {
		return nil, &errorHandler.HttpErr{
			Err:        err,
			StatusCode: http.StatusBadRequest,
		}
	}

	areaResponse = mapper.AreaToAreaResponse(area)

	return areaResponse, nil
}

func (a *AreaService) Delete(id int) error {
	return a.areaRepo.Delete(id)
}

func (a *AreaService) Search(params *filter.AreaFilterParams) (*[]response.Area, *errorHandler.HttpErr) {
	areaResponses := &[]response.Area{}

	areas, err := a.areaRepo.Search(params)
	if err != nil {
		return nil, &errorHandler.HttpErr{
			Err:        err,
			StatusCode: http.StatusBadRequest,
		}
	}

	areaResponses = mapper.AreasToAreaResponses(areas)

	return areaResponses, nil
}
