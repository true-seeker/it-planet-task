package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/LocationValidator"
	"net/http"
)

type LocationHandler struct {
	service service.Location
}

func NewLocationHandler(service service.Location) *LocationHandler {
	return &LocationHandler{service: service}
}

func (l *LocationHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "id must be greater than 0")
		return
	}

	location, err := l.service.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	}
	c.JSON(http.StatusOK, location)
}

func (l *LocationHandler) Create(c *gin.Context) {
	newLocation := &entity.Location{}
	err := c.BindJSON(&newLocation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	httpErr := LocationValidator.ValidateLocation(newLocation)
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}

	duplicateLocation, err := l.service.GetByCords(newLocation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	if duplicateLocation.Id != 0 {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	location, err := l.service.Create(newLocation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, location)
}

func (l *LocationHandler) Update(c *gin.Context) {

}
func (l *LocationHandler) Delete(c *gin.Context) {

}
