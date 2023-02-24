package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"it-planet-task/internal/app/service"
	"it-planet-task/pkg/converter"
	"net/http"
)

type LocationHandler struct {
	service service.Location
}

func NewLocationHandler(service service.Location) *LocationHandler {
	return &LocationHandler{service: service}
}

func (l *LocationHandler) Get(c *gin.Context) {
	id, err := converter.StringToInt(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	location, err := l.service.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)

		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		}
	}
	c.JSON(http.StatusOK, location)
}
