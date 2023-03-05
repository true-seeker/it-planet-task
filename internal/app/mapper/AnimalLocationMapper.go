package mapper

import (
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
)

func AnimalLocationToAnimalLocationResponse(animalLocation *entity.AnimalLocation) *response.AnimalLocation {
	r := &response.AnimalLocation{
		Id:                           animalLocation.Id,
		DateTimeOfVisitLocationPoint: animalLocation.DateTimeOfVisitLocationPoint,
		LocationPointId:              animalLocation.LocationPointId,
	}

	return r
}

func AnimalLocationsToAnimalLocationResponses(animalLocations *[]entity.AnimalLocation) *[]response.AnimalLocation {
	rs := make([]response.AnimalLocation, 0)

	for _, animalLocation := range *animalLocations {
		r := AnimalLocationToAnimalLocationResponse(&animalLocation)
		rs = append(rs, *r)
	}

	return &rs
}
