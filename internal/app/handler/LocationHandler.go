package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/LocationValidator"
	"net/http"
)

// LocationHandler Обработчик запросов для сущности "Локация"
type LocationHandler struct {
	service       service.Location
	animalService service.Animal
}

func NewLocationHandler(service service.Location, animalService service.Animal) *LocationHandler {
	return &LocationHandler{service: service, animalService: animalService}
}

func (l *LocationHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	location, httpErr := l.service.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	c.JSON(http.StatusOK, location)
}

func (l *LocationHandler) Create(c *gin.Context) {
	newLocation := &entity.Location{}
	err := c.BindJSON(&newLocation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	httpErr := LocationValidator.ValidateLocation(newLocation)
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}

	duplicateLocation, err := l.service.GetByCords(newLocation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if duplicateLocation.Id != 0 {
		c.AbortWithStatusJSON(http.StatusConflict, "Location already exists")
		return
	}

	location, err := l.service.Create(newLocation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, location)
}

func (l *LocationHandler) Update(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	oldLocation, httpErr := l.service.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	newLocation := &entity.Location{}
	err := c.BindJSON(&newLocation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	httpErr = LocationValidator.ValidateLocation(newLocation)
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}

	duplicateLocation, err := l.service.GetByCords(newLocation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if duplicateLocation.Id != 0 && oldLocation.Id != duplicateLocation.Id {
		c.AbortWithStatusJSON(http.StatusConflict, "Location already exists")
		return
	}
	newLocation.Id = id
	location, err := l.service.Update(newLocation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, location)
}

func (l *LocationHandler) Delete(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	_, httpErr = l.service.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animals, err := l.animalService.GetAnimalsByLocationId(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if len(*animals) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("There are animals with loocation id %d", id))
		return
	}

	err = l.service.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
