package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"net/http"
)

type AnimalTypeHandler struct {
	service service.AnimalType
}

func NewAnimalTypeHandler(service service.AnimalType) *AnimalTypeHandler {
	return &AnimalTypeHandler{service: service}
}

func (a *AnimalTypeHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "id must be greater than 0")
		return
	}

	animalType, err := a.service.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	}
	c.JSON(http.StatusOK, animalType)
}

func (a *AnimalTypeHandler) Create(c *gin.Context) {
}

func (a *AnimalTypeHandler) Update(c *gin.Context) {

}

func (a *AnimalTypeHandler) Delete(c *gin.Context) {

}
