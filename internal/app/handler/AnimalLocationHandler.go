package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/model/input"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AnimalValidator"
	"net/http"
)

// AnimalLocationHandler Обработчик запросов для сущности "Точка локации живтоного"
type AnimalLocationHandler struct {
	animalLocationService service.AnimalLocation
	animalService         service.Animal
	locationService       service.Location
}

func NewAnimalLocationHandler(animalLocationService service.AnimalLocation, animalService service.Animal, locationService service.Location) *AnimalLocationHandler {
	return &AnimalLocationHandler{animalLocationService: animalLocationService, animalService: animalService, locationService: locationService}
}

func (a *AnimalLocationHandler) GetAnimalLocations(c *gin.Context) {
	animalId, httpErr := validator.ValidateAndReturnId(c.Param("id"), "animalId")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	params, httpErr := filter.NewAnimalLocationFilterParams(c.Request.URL.Query())
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animal, httpErr := a.animalLocationService.GetAnimalLocations(animalId, params)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	c.JSON(http.StatusOK, animal)
}

func (a *AnimalLocationHandler) AddAnimalLocationPoint(c *gin.Context) {
	animalId, httpErr := validator.ValidateAndReturnId(c.Param("id"), "animalId")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	pointId, httpErr := validator.ValidateAndReturnId(c.Param("pointId"), "pointId")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalResponse, httpErr := a.animalService.Get(animalId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	if animalResponse.LifeStatus == AnimalValidator.Dead {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Animal is dead")
		return
	}

	pointResponse, httpErr := a.locationService.Get(pointId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	if len(animalResponse.VisitedLocationsId) == 0 {
		if pointResponse.Id == animalResponse.ChippingLocationId {
			c.AbortWithStatusJSON(http.StatusBadRequest, "animal is in chipping location and hasn't move yet")
			return
		}
	} else {
		if animalResponse.VisitedLocationsId[len(animalResponse.VisitedLocationsId)-1] == pointId {
			c.AbortWithStatusJSON(http.StatusBadRequest, "cat add location point where animal currently is")
			return
		}
	}

	params, httpErr := filter.NewAnimalLocationFilterParams(c.Request.URL.Query())
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	visitedLocations, httpErr := a.animalLocationService.GetAnimalLocations(animalId, params)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	if len(*visitedLocations) > 0 {
		if (*visitedLocations)[len(*visitedLocations)-1].LocationPointId == pointId {
			c.AbortWithStatusJSON(http.StatusBadRequest, "duplicated location")
			return
		}
	}

	animalLocationResponse, err := a.animalLocationService.AddAnimalLocationPoint(animalId, pointId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, animalLocationResponse)
}

func (a *AnimalLocationHandler) EditAnimalLocationPoint(c *gin.Context) {
	animalId, httpErr := validator.ValidateAndReturnId(c.Param("id"), "animalId")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalLocationPointUpdateInput := &input.AnimalLocationPointUpdate{}
	err := c.BindJSON(&animalLocationPointUpdateInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	httpErr = AnimalValidator.ValidateAnimalLocationPointUpdate(animalLocationPointUpdateInput)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalResponse, httpErr := a.animalService.Get(animalId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalLocationsMap := make(map[int]bool)
	for _, animalLocationId := range animalResponse.VisitedLocationsId {
		animalLocationsMap[animalLocationId] = true
	}
	if !animalLocationsMap[*animalLocationPointUpdateInput.VisitedLocationPointId] {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Animal does not have location with id %d", *animalLocationPointUpdateInput.VisitedLocationPointId))
		return
	}

	_, httpErr = a.locationService.Get(*animalLocationPointUpdateInput.LocationPointId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	params, httpErr := filter.NewAnimalLocationFilterParams(c.Request.URL.Query())
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	visitedLocations, httpErr := a.animalLocationService.GetAnimalLocations(animalId, params)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	firstVisitedLocation := (*visitedLocations)[0]
	lastVisitedLocation := (*visitedLocations)[len(*visitedLocations)-1]
	if *animalLocationPointUpdateInput.LocationPointId == animalResponse.ChippingLocationId && firstVisitedLocation.Id == *animalLocationPointUpdateInput.VisitedLocationPointId {
		c.AbortWithStatusJSON(http.StatusBadRequest, "cant set first animal location to chipping location point")
		return
	}
	if lastVisitedLocation.LocationPointId == *animalLocationPointUpdateInput.LocationPointId {
		c.AbortWithStatusJSON(http.StatusBadRequest, "cant set animal location point to last location point")
		return
	}

	for index, visitedLocation := range *visitedLocations {
		if visitedLocation.Id == *animalLocationPointUpdateInput.VisitedLocationPointId {
			if index > 0 {
				previousLocationPoint := (*visitedLocations)[index-1]
				if previousLocationPoint.LocationPointId == *animalLocationPointUpdateInput.LocationPointId {
					c.AbortWithStatusJSON(http.StatusBadRequest, "cant set location point to previous location point")
					return
				}
			}
			if index < len(*visitedLocations)-1 {
				nextLocationPoint := (*visitedLocations)[index+1]
				if nextLocationPoint.LocationPointId == *animalLocationPointUpdateInput.LocationPointId {
					c.AbortWithStatusJSON(http.StatusBadRequest, "cant set location point to next location point")
					return
				}
			}
			break
		}
	}

	animalLocationResponse, err := a.animalLocationService.EditAnimalLocationPoint(*animalLocationPointUpdateInput.VisitedLocationPointId, *animalLocationPointUpdateInput.LocationPointId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, animalLocationResponse)

}

func (a *AnimalLocationHandler) DeleteAnimalLocationPoint(c *gin.Context) {
	animalId, httpErr := validator.ValidateAndReturnId(c.Param("id"), "animalId")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	visitedPointId, httpErr := validator.ValidateAndReturnId(c.Param("visitedPointId"), "visitedPointId")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalResponse, httpErr := a.animalService.Get(animalId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalLocationsMap := make(map[int]bool)
	for _, animalLocationId := range animalResponse.VisitedLocationsId {
		animalLocationsMap[animalLocationId] = true
	}
	if !animalLocationsMap[visitedPointId] {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Animal does not have location with id %d", visitedPointId))
		return
	}

	params, httpErr := filter.NewAnimalLocationFilterParams(c.Request.URL.Query())
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	visitedLocations, httpErr := a.animalLocationService.GetAnimalLocations(animalId, params)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	if len(*visitedLocations) >= 2 {
		if (*visitedLocations)[0].Id == visitedPointId && (*visitedLocations)[1].LocationPointId == animalResponse.ChippingLocationId {
			err := a.animalLocationService.DeleteAnimalLocationPoint((*visitedLocations)[1].Id)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
				return
			}
		}
	}

	err := a.animalLocationService.DeleteAnimalLocationPoint(visitedPointId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
