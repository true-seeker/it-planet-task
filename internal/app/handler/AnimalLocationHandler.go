package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"it-planet-task/internal/app/model/input"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AnimalValidator"
	"net/http"
)

type AnimalLocationHandler struct {
	service         service.AnimalLocation
	animalService   service.Animal
	locationService service.Location
}

func NewAnimalLocationHandler(service service.AnimalLocation, animalService service.Animal, locationService service.Location) *AnimalLocationHandler {
	return &AnimalLocationHandler{service: service, animalService: animalService, locationService: locationService}
}

func (a *AnimalLocationHandler) GetAnimalLocations(c *gin.Context) {
	animalId, httpErr := validator.ValidateAndReturnId(c.Param("id"), "animalId")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animal, httpErr := a.service.GetAnimalLocations(animalId)
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
	// todo Животное находится в точке чипирования и никуда не перемещалось

	_, httpErr = a.locationService.Get(pointId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalLocationResponse, err := a.service.AddAnimalLocationPoint(animalId, pointId)
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
	// TODO validate

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

	// TODO AnimalLocation exists
	_, httpErr = a.locationService.Get(*animalLocationPointUpdateInput.LocationPointId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalLocationResponse, err := a.service.EditAnimalLocationPoint(*animalLocationPointUpdateInput.VisitedLocationPointId, *animalLocationPointUpdateInput.LocationPointId)
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

	err := a.service.DeleteAnimalLocationPoint(visitedPointId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.Status(http.StatusOK)

}
