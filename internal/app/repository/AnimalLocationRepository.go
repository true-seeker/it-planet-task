package repository

import (
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

	animalsHavePreviousPointMap := make(map[int]bool)
	animalsMap := make(map[int]entity.Animal)
	err := a.Db.Raw(`
     WITH 
     q1 as (SELECT a1.id animal_id,
                   a1.chipping_location_id,
                   al1.date_time_of_visit_location_point,
                   l1.latitude,
                   l1.longitude,
                   false is_previous
            FROM animal_locations al1
                     JOIN locations l1 on l1.id = al1.location_point_id
                     JOIN animals a1 on a1.id = al1.animal_id
            WHERE date_time_of_visit_location_point >= ?
              AND date_time_of_visit_location_point <= ?),

     q2 as (SELECT a2.id                                        animal_id2,
                   min(q1.date_time_of_visit_location_point) as date_time_of_visit_location_point
            FROM q1
                     JOIN animals a2 on a2.id = q1.animal_id
            GROUP BY (animal_id2)),

     q3 as (SELECT a3.id                                      as animal_id3,
                   max(al3.date_time_of_visit_location_point) as date_time_of_visit_location_point,
                   true                                          is_previous
            FROM animal_locations al3
                     JOIN locations l3 on l3.id = al3.location_point_id
                     JOIN animals a3 on a3.id = al3.animal_id
                     JOIN q2 on q2.animal_id2 = a3.id
            WHERE al3.date_time_of_visit_location_point < q2.date_time_of_visit_location_point
            GROUP BY (animal_id3))

	 SELECT *
	 FROM q1
	 UNION ALL
	 SELECT a4.id animal_id4,
		   a4.chipping_location_id,
		   al4.date_time_of_visit_location_point,
		   l4.latitude,
		   l4.longitude,
		   q3.is_previous
	 FROM animal_locations al4
			 JOIN locations l4 on l4.id = al4.location_point_id
			 JOIN animals a4 on a4.id = al4.animal_id
			 JOIN q3 on q3.animal_id3 = a4.id
	 WHERE al4.date_time_of_visit_location_point = q3.date_time_of_visit_location_point 
     ORDER BY animal_id, date_time_of_visit_location_point;`, params.StartDateTime, params.EndDateTime).
		Preload("ChippingLocation").
		Scan(&animalLocationsForAreaAnalyticsDTO).
		Error

	animals, err := a.animalRepository.Search(nil)
	if err != nil {
		return nil, err
	}

	for _, animal := range *animals {
		animalsHavePreviousPointMap[animal.Id] = false
		animalsMap[animal.Id] = animal
	}

	for _, animalLocationForAreaAnalyticsDTO := range animalLocationsForAreaAnalyticsDTO {
		animalLocationForAreaAnalytics := entity.AnimalLocationForAreaAnalytics{
			DateTimeOfVisitLocationPoint: animalLocationForAreaAnalyticsDTO.DateTimeOfVisitLocationPoint,
			Location: entity.Location{
				Latitude:  animalLocationForAreaAnalyticsDTO.Latitude,
				Longitude: animalLocationForAreaAnalyticsDTO.Longitude,
			},
			Animal:     animalsMap[animalLocationForAreaAnalyticsDTO.AnimalId],
			IsPrevious: animalLocationForAreaAnalyticsDTO.IsPrevious,
		}
		animalLocationsForAreaAnalytics = append(animalLocationsForAreaAnalytics, animalLocationForAreaAnalytics)
		if animalLocationForAreaAnalyticsDTO.IsPrevious {
			animalsHavePreviousPointMap[animalLocationForAreaAnalyticsDTO.AnimalId] = animalLocationForAreaAnalyticsDTO.IsPrevious
		}
	}

	for animalId, hasPreviousPoint := range animalsHavePreviousPointMap {
		if !hasPreviousPoint {
			animalLocationForAreaAnalytics := entity.AnimalLocationForAreaAnalytics{
				DateTimeOfVisitLocationPoint: animalsMap[animalId].ChippingDateTime,
				Location: entity.Location{
					Latitude:  animalsMap[animalId].ChippingLocation.Latitude,
					Longitude: animalsMap[animalId].ChippingLocation.Longitude,
				},
				Animal:     animalsMap[animalId],
				IsPrevious: true,
			}
			animalLocationsForAreaAnalytics = append([]entity.AnimalLocationForAreaAnalytics{animalLocationForAreaAnalytics}, animalLocationsForAreaAnalytics...)
		}
	}

	return &animalLocationsForAreaAnalytics, nil
}

func (a *AnimalLocationRepository) AddAnimalLocationPoint(newAnimalLocation *entity.AnimalLocation) (*entity.AnimalLocation, error) {
	a.Db.Save(newAnimalLocation)

	return a.Get(newAnimalLocation.Id)
}

func (a *AnimalLocationRepository) EditAnimalLocationPoint(visitedLocationPointId int, locationPointId int) (*entity.AnimalLocation, error) {
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
