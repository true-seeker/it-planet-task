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
	"it-planet-task/internal/app/service/geometry"
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
	Analytics(areaId int, params *filter.AreaAnalyticsFilterParams) (*response.AreaAnalytics, *errorHandler.HttpErr)
}

type AreaService struct {
	areaRepo              repository.Area
	animalLocationService AnimalLocation
	geometryService       geometry.Geometry
}

func NewAreaService(areaRepo repository.Area, animalLocationService AnimalLocation, geometryService geometry.Geometry) Area {
	return &AreaService{areaRepo: areaRepo, animalLocationService: animalLocationService, geometryService: geometryService}
}

func (a *AreaService) Get(id int) (*response.Area, *errorHandler.HttpErr) {
	areaResponse := &response.Area{}

	area, err := a.areaRepo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorHandler.NewHttpErr(fmt.Sprintf("Area with id %d does not exists", id), http.StatusNotFound)
		} else {
			return nil, errorHandler.NewHttpErr(err.Error(), http.StatusBadRequest)
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
			return nil, errorHandler.NewHttpErr(err.Error(), http.StatusBadRequest)
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
		return nil, errorHandler.NewHttpErr(err.Error(), http.StatusBadRequest)
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
			return nil, errorHandler.NewHttpErr(err.Error(), http.StatusBadRequest)
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
		return nil, errorHandler.NewHttpErr(err.Error(), http.StatusBadRequest)
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
		return nil, errorHandler.NewHttpErr(err.Error(), http.StatusBadRequest)
	}

	areaResponses = mapper.AreasToAreaResponses(areas)

	return areaResponses, nil
}

func setTypeMap(mp map[int]map[int]bool, animalTypeId int, animalId int, value bool) {
	_, ok := mp[animalTypeId]
	if !ok {
		mp[animalTypeId] = make(map[int]bool)
	}
	mp[animalTypeId][animalId] = value
}

func countTrueEntities(counter *int, countable map[int]bool) {
	for _, value := range countable {
		if value {
			*counter++
		}
	}
}

func getLenOfMapIfExists(value map[int]bool) int {
	if value != nil {
		return len(value)
	}
	return 0
}

func (a *AreaService) Analytics(areaId int, params *filter.AreaAnalyticsFilterParams) (*response.AreaAnalytics, *errorHandler.HttpErr) {
	areaAnalyticsResponse := response.AreaAnalytics{
		AnimalsAnalytics: []response.AnimalAnalytics{},
	}
	uniqueTypeExits := make(map[int]map[int]bool)
	uniqueTypeEntries := make(map[int]map[int]bool)
	uniqueAreaExits := make(map[int]bool)
	uniqueAreaEntries := make(map[int]bool)
	isTypeInsideArea := make(map[int]map[int]bool)
	isAnimalInsideArea := make(map[int]bool)
	animalTypes := make(map[int]string)

	areaResponse, httpErr := a.Get(areaId)
	if httpErr != nil {
		return nil, httpErr
	}

	area := mapper.AreaResponseToArea(areaResponse)

	points, httpErr := a.animalLocationService.SearchForAreaAnalytics(params)
	if httpErr != nil {
		return nil, httpErr
	}

	// проходим по каждой найденной точке локации животных
	for _, point := range *points {
		if point.IsPrevious {
			// если точка отмечена как предыдущая, то определяем входит ли она в зону и проставляем соответствующие флаги
			// это нужно для определения вошло ли животное в зону из точки, которая не удовлетворяет параметрам запроса
			isAnimalInsideArea[point.Animal.Id] = a.geometryService.IsPointInsideArea(mapper.LocationToAreaPoint(&point.Location), area, true)
			uniqueAreaExits[point.Animal.Id] = false
			uniqueAreaEntries[point.Animal.Id] = false
			if isAnimalInsideArea[point.Animal.Id] {
				for _, animalType := range point.Animal.AnimalTypes {
					animalTypes[animalType.Id] = animalType.Type
					setTypeMap(isTypeInsideArea, animalType.Id, point.Animal.Id, true)
				}
			}
		} else {
			if a.geometryService.IsPointInsideArea(mapper.LocationToAreaPoint(&point.Location), area, true) {
				if !isAnimalInsideArea[point.Animal.Id] {
					// если очередная точка в зоне, но до этого животное было вне
					for _, animalType := range point.Animal.AnimalTypes {
						animalTypes[animalType.Id] = animalType.Type
						setTypeMap(uniqueTypeEntries, animalType.Id, point.Animal.Id, true)
						setTypeMap(isTypeInsideArea, animalType.Id, point.Animal.Id, true)
					}
					isAnimalInsideArea[point.Animal.Id] = true
					uniqueAreaEntries[point.Animal.Id] = true
				}

			} else {
				if isAnimalInsideArea[point.Animal.Id] {
					// если очередная точка не в зоне, но до этого животное было в ней
					for _, animalType := range point.Animal.AnimalTypes {
						animalTypes[animalType.Id] = animalType.Type
						setTypeMap(uniqueTypeExits, animalType.Id, point.Animal.Id, true)
						setTypeMap(isTypeInsideArea, animalType.Id, point.Animal.Id, false)
					}
					isAnimalInsideArea[point.Animal.Id] = false
					uniqueAreaExits[point.Animal.Id] = true
				}
			}
		}
	}

	// подсчёт статистики для каждого типа
	for animalTypeId, animalType := range animalTypes {
		animalAnalytics := response.AnimalAnalytics{
			AnimalType:     animalType,
			AnimalTypeId:   animalTypeId,
			AnimalsArrived: getLenOfMapIfExists(uniqueTypeEntries[animalTypeId]),
			AnimalsGone:    getLenOfMapIfExists(uniqueTypeExits[animalTypeId]),
		}
		countTrueEntities(&animalAnalytics.QuantityAnimals, isTypeInsideArea[animalTypeId])
		areaAnalyticsResponse.AnimalsAnalytics = append(areaAnalyticsResponse.AnimalsAnalytics, animalAnalytics)
	}

	// подсчёт общей статистики
	countTrueEntities(&areaAnalyticsResponse.TotalQuantityAnimals, isAnimalInsideArea)
	countTrueEntities(&areaAnalyticsResponse.TotalAnimalsArrived, uniqueAreaEntries)
	countTrueEntities(&areaAnalyticsResponse.TotalAnimalsGone, uniqueAreaExits)
	return &areaAnalyticsResponse, nil
}
