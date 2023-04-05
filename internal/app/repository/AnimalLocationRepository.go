package repository

import (
	"fmt"
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/pkg/paginator"
)

type AnimalLocation interface {
	GetAnimalLocations(animalId int, params *filter.AnimalLocationFilterParams) (*[]entity.AnimalLocation, error)
	AddAnimalLocationPoint(newAnimalLocation *entity.AnimalLocation) (*entity.AnimalLocation, error)
	EditAnimalLocationPoint(visitedLocationPointId int, locationPointId int) (*entity.AnimalLocation, error)
	DeleteAnimalLocationPoint(id int) error
	Get(id int) (*entity.AnimalLocation, error)
	SearchForAreaAnalytics(params *filter.AnimalLocationFilterParams) (*[]entity.AnimalLocationForAreaAnalytics, error)
}

type AnimalLocationRepository struct {
	Db               *gorm.DB
	animalRepository Animal
}

func NewAnimalLocationRepository(db *gorm.DB, animalRepository Animal) AnimalLocation {
	return &AnimalLocationRepository{Db: db, animalRepository: animalRepository}
}

func (a *AnimalLocationRepository) GetAnimalLocations(animalId int, params *filter.AnimalLocationFilterParams) (*[]entity.AnimalLocation, error) {
	var animalLocations []entity.AnimalLocation

	err := a.Db.Where("animal_id = ?", animalId).
		Scopes(paginator.Paginate(params),
			filter.AnimalLocationFilter(params)).
		Find(&animalLocations).Error

	if err != nil {
		return nil, err
	}

	return &animalLocations, nil
}

func (a *AnimalLocationRepository) SearchForAreaAnalytics(params *filter.AnimalLocationFilterParams) (*[]entity.AnimalLocationForAreaAnalytics, error) {
	var animalLocationsForAreaAnalytics []entity.AnimalLocationForAreaAnalytics
	var animalLocationsForAreaAnalyticsDTO []response.AnimalLocationForAreaAnalyticsDTO
	var animalIds []int
	animalsForAreaAnalyticsMap := make(map[int]entity.Animal)
	err := a.Db.
		Table("animal_locations as al").
		Select("a.id as animal_id, al.date_time_of_visit_location_point, l.latitude, l.longitude").
		Joins("JOIN locations l on l.id = al.location_point_id").
		Joins("JOIN animals a on a.id = al.animal_id").
		Scopes(filter.AreaAnalyticsFilter(params)).
		Scan(&animalLocationsForAreaAnalyticsDTO).
		Error
	if err != nil {
		return nil, err
	}

	for _, animalLocationForAreaAnalyticsDTO := range animalLocationsForAreaAnalyticsDTO {
		animalIds = append(animalIds, animalLocationForAreaAnalyticsDTO.AnimalId)
	}

	animals, err := a.animalRepository.GetByIds(&animalIds)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	for _, animal := range *animals {
		animalsForAreaAnalyticsMap[animal.Id] = animal
	}

	for _, animalLocationForAreaAnalyticsDTO := range animalLocationsForAreaAnalyticsDTO {
		animalLocationForAreaAnalytics := entity.AnimalLocationForAreaAnalytics{
			DateTimeOfVisitLocationPoint: animalLocationForAreaAnalyticsDTO.DateTimeOfVisitLocationPoint,
			Location: entity.Location{
				Latitude:  animalLocationForAreaAnalyticsDTO.Latitude,
				Longitude: animalLocationForAreaAnalyticsDTO.Longitude,
			},
			Animal: animalsForAreaAnalyticsMap[animalLocationForAreaAnalyticsDTO.AnimalId],
		}
		animalLocationsForAreaAnalytics = append(animalLocationsForAreaAnalytics, animalLocationForAreaAnalytics)
	}

	return &animalLocationsForAreaAnalytics, nil
}

func (a *AnimalLocationRepository) AddAnimalLocationPoint(newAnimalLocation *entity.AnimalLocation) (*entity.AnimalLocation, error) {
	a.Db.Save(newAnimalLocation)

	return a.Get(newAnimalLocation.Id)
}

func (a *AnimalLocationRepository) EditAnimalLocationPoint(visitedLocationPointId int, locationPointId int) (*entity.AnimalLocation, error) {
	fmt.Println(locationPointId, visitedLocationPointId)
	a.Db.Exec("UPDATE animal_locations SET location_point_id = ? WHERE id = ?", locationPointId, visitedLocationPointId)

	return a.Get(visitedLocationPointId)
}

func (a *AnimalLocationRepository) DeleteAnimalLocationPoint(id int) error {
	err := a.Db.Delete(&entity.AnimalLocation{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *AnimalLocationRepository) Get(id int) (*entity.AnimalLocation, error) {
	var animalLocation entity.AnimalLocation
	err := a.Db.First(&animalLocation, id).Error

	if err != nil {
		return nil, err
	}

	return &animalLocation, nil
}
