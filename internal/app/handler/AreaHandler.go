package handler

import (
	"github.com/gin-gonic/gin"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/repository"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AreaValidator"
	"net/http"
)

type AreaHandler struct {
	areaService service.Area
	areaRepo    repository.Area
}

func NewAreaHandler(areaService service.Area, areaRepo repository.Area) *AreaHandler {
	return &AreaHandler{areaService: areaService, areaRepo: areaRepo}
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
	newArea := &entity.Area{}
	err := c.BindJSON(&newArea)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	httpErr := AreaValidator.ValidateArea(newArea)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	area, httpErr := a.areaService.Create(newArea)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	c.JSON(http.StatusCreated, area)
}

func (a *AreaHandler) Update(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	_, httpErr = a.areaService.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	newArea := &entity.Area{}
	err := c.BindJSON(&newArea)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	httpErr = AreaValidator.ValidateArea(newArea)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	newArea.Id = id

	area, httpErr := a.areaService.Update(newArea)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	c.JSON(http.StatusOK, area)
}

func (a *AreaHandler) Delete(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	err := a.areaService.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
