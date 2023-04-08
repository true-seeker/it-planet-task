package AnimalValidator

import (
	"fmt"
	"it-planet-task/internal/app/model/input"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

const (
	Alive = "ALIVE"
	Dead  = "DEAD"
)
const (
	Male   = "MALE"
	Female = "FEMALE"
	Other  = "OTHER"
)

func ValidateLifeStatus(lifeStatus string) *errorHandler.HttpErr {
	if lifeStatus != Alive && lifeStatus != Dead {
		return errorHandler.NewHttpErr(fmt.Sprintf("lifeStatus must be in [%s, %s]", Alive, Dead), http.StatusBadRequest)
	}
	return nil
}

func ValidateGender(gender string) *errorHandler.HttpErr {
	if gender != Male && gender != Female && gender != Other {
		return errorHandler.NewHttpErr(fmt.Sprintf("gender must be in [%s, %s, %s]", Male, Female, Other), http.StatusBadRequest)
	}
	return nil
}

func validateAnimalInput(input *input.Animal) *errorHandler.HttpErr {
	if input.Weight == nil {
		return errorHandler.NewHttpErr("weight is missing", http.StatusBadRequest)
	}
	if *input.Weight <= 0 {
		return errorHandler.NewHttpErr("weight must be greater than 0", http.StatusBadRequest)
	}

	if input.Length == nil {
		return errorHandler.NewHttpErr("length is missing", http.StatusBadRequest)
	}
	if *input.Length <= 0 {
		return errorHandler.NewHttpErr("length must be greater than 0", http.StatusBadRequest)
	}

	if input.Height == nil {
		return errorHandler.NewHttpErr("height is missing", http.StatusBadRequest)
	}
	if *input.Height <= 0 {
		return errorHandler.NewHttpErr("height must be greater than 0", http.StatusBadRequest)
	}

	if input.Gender == nil {
		return errorHandler.NewHttpErr("gender is missing", http.StatusBadRequest)
	}
	err := ValidateGender(*input.Gender)
	if err != nil {
		return err
	}

	if input.ChipperId == nil {
		return errorHandler.NewHttpErr("chipper_id is missing", http.StatusBadRequest)
	}
	if *input.ChipperId <= 0 {
		return errorHandler.NewHttpErr("chipper_id must be greater than 0", http.StatusBadRequest)
	}

	if input.ChippingLocationId == nil {
		return errorHandler.NewHttpErr("chipping_location_id is missing", http.StatusBadRequest)
	}
	if *input.ChippingLocationId <= 0 {
		return errorHandler.NewHttpErr("chipping_location_id must be greater than 0", http.StatusBadRequest)
	}
	return nil
}

func ValidateAnimalCreateInput(input *input.Animal) *errorHandler.HttpErr {
	if input.AnimalTypeIds == nil || len(input.AnimalTypeIds) == 0 {
		return errorHandler.NewHttpErr("animal types are empty", http.StatusBadRequest)
	}

	animalTypeIds := map[int]bool{}
	for _, animalTypeId := range input.AnimalTypeIds {
		if animalTypeId <= 0 {
			return errorHandler.NewHttpErr("animal type id must be greater than 0", http.StatusBadRequest)
		}
		if animalTypeIds[animalTypeId] {
			return errorHandler.NewHttpErr("duplicated animal type id", http.StatusConflict)
		}
		animalTypeIds[animalTypeId] = true
	}

	httpErr := validateAnimalInput(input)
	if httpErr != nil {
		return httpErr
	}

	return nil
}

func ValidateAnimalUpdateInput(input *input.Animal, oldAnimal *response.Animal) *errorHandler.HttpErr {
	httpErr := validateAnimalInput(input)
	if httpErr != nil {
		return httpErr
	}

	if *input.LifeStatus == Alive && oldAnimal.LifeStatus == Dead {
		return errorHandler.NewHttpErr("cant set status Alive to Dead animal", http.StatusBadRequest)
	}
	return nil
}

func ValidateAnimalTypeUpdateInput(input *input.AnimalTypeUpdate) *errorHandler.HttpErr {
	if input.OldTypeId == nil {
		return errorHandler.NewHttpErr("oldTypeId is missing", http.StatusBadRequest)
	}
	if *input.OldTypeId <= 0 {
		return errorHandler.NewHttpErr("oldTypeId must be greater than 0", http.StatusBadRequest)
	}
	if input.NewTypeId == nil {
		return errorHandler.NewHttpErr("newTypeId is missing", http.StatusBadRequest)
	}
	if *input.NewTypeId <= 0 {
		return errorHandler.NewHttpErr("newTypeId must be greater than 0", http.StatusBadRequest)
	}
	return nil
}

func ValidateAnimalLocationPointUpdate(input *input.AnimalLocationPointUpdate) *errorHandler.HttpErr {
	if input.VisitedLocationPointId == nil {
		return errorHandler.NewHttpErr("visitedLocationPointId is missing", http.StatusBadRequest)
	}
	if *input.VisitedLocationPointId <= 0 {
		return errorHandler.NewHttpErr("visitedLocationPointId must be greater than 0", http.StatusBadRequest)
	}
	if input.LocationPointId == nil {
		return errorHandler.NewHttpErr("locationPointId is missing", http.StatusBadRequest)
	}
	if *input.LocationPointId <= 0 {
		return errorHandler.NewHttpErr("locationPointId must be greater than 0", http.StatusBadRequest)
	}
	return nil
}
