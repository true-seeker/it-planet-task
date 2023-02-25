package mapper

import (
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
)

func AnimalToAnimalResponse(animal *entity.Animal) *response.Animal {
	r := &response.Animal{
		Id:                 animal.Id,
		AnimalTypes:        nil,
		Weight:             animal.Weight,
		Height:             animal.Height,
		Length:             animal.Length,
		Gender:             animal.Gender,
		LifeStatus:         animal.LifeStatus,
		ChippingDateTime:   animal.ChippingDateTime,
		ChipperId:          animal.ChipperId,
		ChippingLocationId: animal.ChippingLocationId,
		VisitedLocations:   nil,
		DeathDateTime:      animal.DeathDateTime,
	}

	for _, visitedLoc := range animal.VisitedLocations {
		r.VisitedLocations = append(r.VisitedLocations, visitedLoc.Id)
	}

	for _, animalType := range animal.AnimalTypes {
		r.AnimalTypes = append(r.AnimalTypes, animalType.Id)
	}

	return r
}

func AnimalsToAnimalResponses(animals *[]entity.Animal) *[]response.Animal {
	var r []response.Animal
	for _, animal := range *animals {
		r = append(r, *AnimalToAnimalResponse(&animal))
	}

	return &r
}
