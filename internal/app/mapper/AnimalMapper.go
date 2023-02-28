package mapper

import (
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/input"
	"it-planet-task/internal/app/model/response"
)

func AnimalToAnimalResponse(animal *entity.Animal) *response.Animal {
	r := &response.Animal{
		Id:                 animal.Id,
		AnimalTypesId:      nil,
		Weight:             animal.Weight,
		Height:             animal.Height,
		Length:             animal.Length,
		Gender:             animal.Gender,
		LifeStatus:         animal.LifeStatus,
		ChippingDateTime:   animal.ChippingDateTime,
		ChipperId:          animal.ChipperId,
		ChippingLocationId: animal.ChippingLocationId,
		VisitedLocationsId: []int{},
		DeathDateTime:      animal.DeathDateTime,
	}

	for _, visitedLoc := range animal.VisitedLocations {
		r.VisitedLocationsId = append(r.VisitedLocationsId, visitedLoc.Id)
	}

	for _, animalType := range animal.AnimalTypes {
		r.AnimalTypesId = append(r.AnimalTypesId, animalType.Id)
	}

	return r
}

func AnimalsToAnimalResponses(animals *[]entity.Animal) *[]response.Animal {
	r := make([]response.Animal, 0)

	for _, animal := range *animals {
		r = append(r, *AnimalToAnimalResponse(&animal))
	}

	return &r
}

func AnimalCreateInputToAnimal(input *input.AnimalCreate) *entity.Animal {
	r := &entity.Animal{
		Weight:             *input.Weight,
		Height:             *input.Height,
		Length:             *input.Length,
		Gender:             *input.Gender,
		ChipperId:          *input.ChipperId,
		ChippingLocationId: *input.ChippingLocationId,
	}

	for _, animalTypeId := range input.AnimalTypeIds {
		r.AnimalTypes = append(r.AnimalTypes, entity.AnimalType{Id: animalTypeId})
	}

	return r
}
