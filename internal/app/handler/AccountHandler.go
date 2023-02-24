package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"it-planet-task/internal/app/service"
	"it-planet-task/pkg/converter"
	"net/http"
)

type AccountHandler struct {
	service service.Account
}

func NewAccountHandler(service service.Account) *AccountHandler {
	return &AccountHandler{service: service}
}

func (a *AccountHandler) Get(c *gin.Context) {
	id, err := converter.StringToInt(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	account, err := a.service.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)

		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		}
	}
	c.JSON(http.StatusOK, account)
}
