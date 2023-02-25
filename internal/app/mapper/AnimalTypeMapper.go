package mapper

import (
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
)

func AnimalTypeToAnimalTypeResponse(animalType *entity.AnimalType) *response.AnimalType {
	r := &response.AnimalType{
		Id:    animalType.Id,
		Title: animalType.Title,
	}

	return r
}
