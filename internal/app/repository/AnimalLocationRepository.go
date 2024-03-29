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
	SearchForAreaAnalytics(params *filter.AreaAnalyticsFilterParams) (*[]entity.AnimalLocationForAreaAnalytics, error)
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

func (a *AnimalLocationRepository) SearchForAreaAnalytics(params *filter.AreaAnalyticsFilterParams) (*[]entity.AnimalLocationForAreaAnalytics, error) {
	var analytics []entity.AnimalLocationForAreaAnalytics
	var animalLocationsAnalyticsDTO []response.AnimalLocationForAreaAnalyticsDTO
	animalsHavePreviousPoint := make(map[int]bool)
	animalsMap := make(map[int]entity.Animal)

	// получение всех точке локации животных, удовлетворяющих диапазону дат params
	// при этом для этих точек ищется предыдущая точка, которая не входит в этот диапазон, если такая есть
	err := a.Db.Raw(`
     WITH 
	-- Получение всех точек, удовлетворяющих диапазону дат
     q1 as (SELECT a1.id animal_id,
                   al1.date_time_of_visit_location_point,
                   l1.latitude,
                   l1.longitude,
                   false is_previous
            FROM animal_locations al1
                     JOIN locations l1 on l1.id = al1.location_point_id
                     JOIN animals a1 on a1.id = al1.animal_id
            WHERE date_time_of_visit_location_point >= ?
              AND date_time_of_visit_location_point <= ?),

	-- выделение из q1 самой ранней даты для каждого животного
     q2 as (SELECT a2.id                                        animal_id2,
                   min(q1.date_time_of_visit_location_point) as date_time_of_visit_location_point
            FROM q1
                     JOIN animals a2 on a2.id = q1.animal_id
            GROUP BY (animal_id2)),

	-- для каждого животного получение одной точки, предыдущей точкам из q2 и помечается как is_previous
     q3 as (SELECT a3.id                                      as animal_id3,
                   max(al3.date_time_of_visit_location_point) as date_time_of_visit_location_point,
                   true                                       as  is_previous
            FROM animal_locations al3
                     JOIN locations l3 on l3.id = al3.location_point_id
                     JOIN animals a3 on a3.id = al3.animal_id
                     JOIN q2 on q2.animal_id2 = a3.id
            WHERE al3.date_time_of_visit_location_point < q2.date_time_of_visit_location_point
            GROUP BY (animal_id3))

	-- объединение q1 и q3
	 SELECT *
	 FROM q1
	 UNION ALL
	 SELECT a4.id animal_id4,
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
		Scan(&animalLocationsAnalyticsDTO).
		Error
	if err != nil {
		return nil, err
	}

	// получаем всех животных, т.к. предыдущий запрос не учитывает точки чипирования
	animals, err := a.animalRepository.Search(nil)
	if err != nil {
		return nil, err
	}

	for _, animal := range *animals {
		animalsHavePreviousPoint[animal.Id] = false
		animalsMap[animal.Id] = animal
	}

	// заполняем сущности для дальнейшей обработки
	for _, analyticsDTO := range animalLocationsAnalyticsDTO {
		location := entity.NewLocation(0, analyticsDTO.Latitude, analyticsDTO.Longitude)
		animalLocationForAreaAnalytics := entity.NewAnimalLocationForAreaAnalytics(analyticsDTO.DateTimeOfVisitLocationPoint, *location,
			animalsMap[analyticsDTO.AnimalId], analyticsDTO.IsPrevious)

		analytics = append(analytics, *animalLocationForAreaAnalytics)
		if analyticsDTO.IsPrevious {
			animalsHavePreviousPoint[analyticsDTO.AnimalId] = analyticsDTO.IsPrevious
		}
	}

	// для всех точек, у которых не нашлось предыдущей, ставим точку чипирования как предыдущую
	for animalId, hasPreviousPoint := range animalsHavePreviousPoint {
		if !hasPreviousPoint {
			location := entity.NewLocation(0, animalsMap[animalId].ChippingLocation.Latitude, animalsMap[animalId].ChippingLocation.Longitude)
			animalLocationAnalytics := entity.NewAnimalLocationForAreaAnalytics(animalsMap[animalId].ChippingDateTime, *location, animalsMap[animalId], true)
			analytics = append([]entity.AnimalLocationForAreaAnalytics{*animalLocationAnalytics}, analytics...)
		}
	}

	return &analytics, nil
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
