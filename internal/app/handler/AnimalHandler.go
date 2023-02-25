package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"it-planet-task/internal/app/service"
	"it-planet-task/pkg/converter"
	"net/http"
)

type AnimalHandler struct {
	service service.Animal
}

func NewAnimalHandler(service service.Animal) *AnimalHandler {
	return &AnimalHandler{service: service}
}

func (a *AnimalHandler) Get(c *gin.Context) {
	id, err := converter.StringToInt(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	animal, err := a.service.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)

		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		}
	}
	c.JSON(http.StatusOK, animal)
}

func (a *AnimalHandler) GetAnimalLocations(c *gin.Context) {
	id, err := converter.StringToInt(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	animal, err := a.service.GetAnimalLocations(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)

		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		}
	}
	c.JSON(http.StatusOK, animal)
}
