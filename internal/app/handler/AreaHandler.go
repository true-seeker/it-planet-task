package handler

import (
	"github.com/gin-gonic/gin"
	"it-planet-task/internal/app/service"
)

type AreaHandler struct {
	areaService service.Area
}

func NewAreaHandler(areaService service.Area) *AreaHandler {
	return &AreaHandler{areaService: areaService}
}

func (a *AreaHandler) Get(c *gin.Context) {

}

func (a *AreaHandler) Create(c *gin.Context) {

}

func (a *AreaHandler) Update(c *gin.Context) {

}

func (a *AreaHandler) Delete(c *gin.Context) {

}
