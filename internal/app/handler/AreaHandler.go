package handler

import (
	"github.com/gin-gonic/gin"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"net/http"
)

type AreaHandler struct {
	areaService service.Area
}

func NewAreaHandler(areaService service.Area) *AreaHandler {
	return &AreaHandler{areaService: areaService}
}

func (a *AreaHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	area, httpErr := a.areaService.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	c.JSON(http.StatusOK, area)
}

func (a *AreaHandler) Create(c *gin.Context) {

}

func (a *AreaHandler) Update(c *gin.Context) {

}

func (a *AreaHandler) Delete(c *gin.Context) {

}
