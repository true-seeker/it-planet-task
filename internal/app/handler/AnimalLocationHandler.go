package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	animalId, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "animalId")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if animalId <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "animalId must be greater than 0")
		return
	}

	animal, err := a.service.GetAnimalLocations(animalId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return

		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	}
	c.JSON(http.StatusOK, animal)
}

func (a *AnimalLocationHandler) AddAnimalLocationPoint(c *gin.Context) {
	animalId, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "animalId")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if animalId <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "animalId must be greater than 0")
		return
	}
	pointId, httpErr := validator.ValidateAndReturnIntField(c.Param("pointId"), "pointId")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if pointId <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "pointId must be greater than 0")
		return
	}

	animalResponse, err := a.animalService.Get(animalId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	if animalResponse.LifeStatus == AnimalValidator.Dead {
		c.AbortWithStatusJSON(http.StatusBadRequest, "animal is dead")
		return
	}
	// todo Животное находится в точке чипирования и никуда не перемещалось

	_, err = a.locationService.Get(pointId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	animalLocationResponse, err := a.service.AddAnimalLocationPoint(animalId, pointId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, animalLocationResponse)
}

func (a *AnimalLocationHandler) EditAnimalLocationPoint(c *gin.Context) {
	animalId, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "animalId")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if animalId <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "animalId must be greater than 0")
		return
	}

	animalLocationPointUpdateInput := &input.AnimalLocationPointUpdate{}
	err := c.BindJSON(&animalLocationPointUpdateInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	httpErr = AnimalValidator.ValidateAnimalLocationPointUpdate(animalLocationPointUpdateInput)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	// TODO validate

	animalResponse, err := a.animalService.Get(animalId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	animalLocationsMap := make(map[int]bool)
	for _, animalLocationId := range animalResponse.VisitedLocationsId {
		animalLocationsMap[animalLocationId] = true
	}
	if !animalLocationsMap[*animalLocationPointUpdateInput.VisitedLocationPointId] {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	// TODO AnimalLocation exists
	_, err = a.locationService.Get(*animalLocationPointUpdateInput.LocationPointId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	animalLocationResponse, err := a.service.EditAnimalLocationPoint(*animalLocationPointUpdateInput.VisitedLocationPointId, *animalLocationPointUpdateInput.LocationPointId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, animalLocationResponse)

}

func (a *AnimalLocationHandler) DeleteAnimalLocationPoint(c *gin.Context) {
	animalId, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "animalId")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if animalId <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "animalId must be greater than 0")
		return
	}

	visitedPointId, httpErr := validator.ValidateAndReturnIntField(c.Param("visitedPointId"), "visitedPointId")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if visitedPointId <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "visitedPointId must be greater than 0")
		return
	}

	animalResponse, err := a.animalService.Get(animalId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	animalLocationsMap := make(map[int]bool)
	for _, animalLocationId := range animalResponse.VisitedLocationsId {
		animalLocationsMap[animalLocationId] = true
	}
	if !animalLocationsMap[visitedPointId] {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	err = a.service.DeleteAnimalLocationPoint(visitedPointId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.Status(http.StatusOK)

}
