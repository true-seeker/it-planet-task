package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
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
	id, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "id must be greater than 0")
		return
	}

	animal, err := a.service.Get(id)
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

func (a *AnimalHandler) GetAnimalLocations(c *gin.Context) {
	id, err := converter.StringToInt(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	animal, err := a.service.GetAnimalLocations(id)
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

func (a *AnimalHandler) Search(c *gin.Context) {
	params, httpErr := filter.NewAnimalFilterParams(c.Request.URL.Query())
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	animals, err := a.service.Search(params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, animals)
}
