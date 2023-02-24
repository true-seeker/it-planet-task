package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"it-planet-task/internal/app/service"
	"it-planet-task/pkg/converter"
	"net/http"
)

type AnimalTypeHandler struct {
	service service.AnimalType
}

func NewAnimalTypeHandler(service service.AnimalType) *AnimalTypeHandler {
	return &AnimalTypeHandler{service: service}
}

func (a *AnimalTypeHandler) Get(c *gin.Context) {
	id, err := converter.StringToInt(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	animalType, err := a.service.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)

		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		}
	}
	c.JSON(http.StatusOK, animalType)
}
