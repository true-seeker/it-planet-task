package AnimalValidator

import (
	"errors"
	"fmt"
	"it-planet-task/internal/app/model/input"
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
		return &errorHandler.HttpErr{
			Err:        errors.New(fmt.Sprintf("lifeStatus must be in [%s, %s]", Alive, Dead)),
			StatusCode: http.StatusBadRequest,
		}
	}
	return nil
}

func ValidateGender(gender string) *errorHandler.HttpErr {
	if gender != Male && gender != Female && gender != Other {
		return &errorHandler.HttpErr{
			Err:        errors.New(fmt.Sprintf("gender must be in [%s, %s, %s]", Male, Female, Other)),
			StatusCode: http.StatusBadRequest,
		}
	}
	return nil
}

func ValidateAnimalCreateInput(input *input.AnimalCreate) *errorHandler.HttpErr {
	if input.AnimalTypeIds == nil || len(input.AnimalTypeIds) == 0 {
		return &errorHandler.HttpErr{
			Err:        errors.New("animal types are empty"),
			StatusCode: http.StatusBadRequest,
		}
	}

	animalTypeIds := map[int]bool{}
	for _, animalTypeId := range input.AnimalTypeIds {
		if animalTypeId <= 0 {
			return &errorHandler.HttpErr{
				Err:        errors.New("animal type id must be greater than 0"),
				StatusCode: http.StatusBadRequest,
			}
		}
		if animalTypeIds[animalTypeId] {
			return &errorHandler.HttpErr{
				Err:        errors.New("duplicated animal type id"),
				StatusCode: http.StatusConflict,
			}
		}
		animalTypeIds[animalTypeId] = true
	}

	if input.Weight == nil {
		return &errorHandler.HttpErr{
			Err:        errors.New("weight is missing"),
			StatusCode: http.StatusBadRequest,
		}
	}
	if *input.Weight <= 0 {
		return &errorHandler.HttpErr{
			Err:        errors.New("weight must be greater than 0"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if input.Length == nil {
		return &errorHandler.HttpErr{
			Err:        errors.New("length is missing"),
			StatusCode: http.StatusBadRequest,
		}
	}
	if *input.Length <= 0 {
		return &errorHandler.HttpErr{
			Err:        errors.New("length must be greater than 0"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if input.Height == nil {
		return &errorHandler.HttpErr{
			Err:        errors.New("height is missing"),
			StatusCode: http.StatusBadRequest,
		}
	}
	if *input.Height <= 0 {
		return &errorHandler.HttpErr{
			Err:        errors.New("height must be greater than 0"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if input.Gender == nil {
		return &errorHandler.HttpErr{
			Err:        errors.New("gender is missing"),
			StatusCode: http.StatusBadRequest,
		}
	}
	err := ValidateGender(*input.Gender)
	if err != nil {
		return err
	}

	if input.ChipperId == nil {
		return &errorHandler.HttpErr{
			Err:        errors.New("chipper_id is missing"),
			StatusCode: http.StatusBadRequest,
		}
	}
	if *input.ChipperId <= 0 {
		return &errorHandler.HttpErr{
			Err:        errors.New("chipper_id must be greater than 0"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if input.ChippingLocationId == nil {
		return &errorHandler.HttpErr{
			Err:        errors.New("chipping_location_id is missing"),
			StatusCode: http.StatusBadRequest,
		}
	}
	if *input.ChippingLocationId <= 0 {
		return &errorHandler.HttpErr{
			Err:        errors.New("chipping_location_id must be greater than 0"),
			StatusCode: http.StatusBadRequest,
		}
	}

	return nil
}
