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

func (a *AreaService) Analytics(areaId int, params *filter.AreaAnalyticsFilterParams) (*response.AreaAnalytics, *errorHandler.HttpErr) {
	var animalAnalyticsResponse []response.AnimalAnalytics
	areaAnalyticsResponse := response.AreaAnalytics{
		AnimalAnalytics: []response.AnimalAnalytics{},
	}
	uniqueAreaExits := make(map[int]map[int]bool)
	uniqueAreaEntries := make(map[int]map[int]bool)
	animalsInsideArea := make(map[int]map[int]bool)
	animalTypes := make(map[int]string)

	areaResponse, httpErr := a.Get(areaId)
	if httpErr != nil {
		return nil, httpErr
	}
	area := mapper.AreaResponseToArea(areaResponse)

	animalLocationParams := &filter.AnimalLocationFilterParams{
		StartDateTime: params.StartDateTime,
		EndDateTime:   params.EndDateTime,
	}
	animalLocationsForAreaAnalytics, httpErr := a.animalLocationService.SearchForAreaAnalytics(animalLocationParams)
	if httpErr != nil {
		return nil, httpErr
	}
	// TODO chipping location
	for _, animalLocationForAreaAnalytics := range *animalLocationsForAreaAnalytics {
		fmt.Println(animalLocationForAreaAnalytics.Animal.Id, animalLocationForAreaAnalytics.DateTimeOfVisitLocationPoint)
		if a.geometryService.IsPointInsideArea(mapper.LocationToAreaPoint(&animalLocationForAreaAnalytics.Location), area, true) {
			for _, animalType := range animalLocationForAreaAnalytics.Animal.AnimalTypes {
				animalTypes[animalType.Id] = animalType.Type
				_, ok := uniqueAreaEntries[animalType.Id]
				if !ok {
					uniqueAreaEntries[animalType.Id] = make(map[int]bool)
				}
				uniqueAreaEntries[animalType.Id][animalLocationForAreaAnalytics.Animal.Id] = true

				_, ok = animalsInsideArea[animalType.Id]
				if !ok {
					animalsInsideArea[animalType.Id] = make(map[int]bool)
				}
				animalsInsideArea[animalType.Id][animalLocationForAreaAnalytics.Animal.Id] = true
			}
		} else {
			for animalTypeId := range uniqueAreaEntries {
				for animalId := range uniqueAreaEntries[animalTypeId] {
					if animalId == animalLocationForAreaAnalytics.Animal.Id {
						_, ok := uniqueAreaExits[animalTypeId]
						if !ok {
							uniqueAreaExits[animalTypeId] = make(map[int]bool)
						}
						uniqueAreaExits[animalTypeId][animalLocationForAreaAnalytics.Animal.Id] = true
						animalsInsideArea[animalTypeId][animalId] = false
					}
				}
			}
		}
	}

	for animalTypeId := range uniqueAreaEntries {
		animalAnalytics := response.AnimalAnalytics{
			AnimalType:   animalTypes[animalTypeId],
			AnimalTypeId: animalTypeId,
		}
		for range uniqueAreaEntries[animalTypeId] {
			animalAnalytics.AnimalsArrived++
		}
		for range uniqueAreaExits[animalTypeId] {
			animalAnalytics.AnimalsGone++
		}
		for _, isAnimalInsideArea := range animalsInsideArea[animalTypeId] {
			if isAnimalInsideArea {
				animalAnalytics.QuantityAnimals++
			}
		}
		animalAnalyticsResponse = append(animalAnalyticsResponse, animalAnalytics)
	}

	for _, animalAnalytics := range animalAnalyticsResponse {
		areaAnalyticsResponse.AnimalAnalytics = append(areaAnalyticsResponse.AnimalAnalytics, animalAnalytics)
		areaAnalyticsResponse.TotalQuantityAnimals += animalAnalytics.QuantityAnimals
		areaAnalyticsResponse.TotalAnimalsArrived += animalAnalytics.AnimalsArrived
		areaAnalyticsResponse.TotalAnimalsGone += animalAnalytics.AnimalsGone
	}

	return &areaAnalyticsResponse, nil
}
