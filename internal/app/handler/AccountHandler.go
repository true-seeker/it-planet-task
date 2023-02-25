package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"net/http"
)

type AccountHandler struct {
	service service.Account
}

func NewAccountHandler(service service.Account) *AccountHandler {
	return &AccountHandler{service: service}
}

func (a *AccountHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "id must be greater than 0")
		return
	}

	account, err := a.service.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return

		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	}
	c.JSON(http.StatusOK, account)
}

func (a *AccountHandler) Search(c *gin.Context) {
	params, httpErr := filter.NewAccountFilterParams(c.Request.URL.Query())
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	accounts, err := a.service.Search(params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, accounts)
}
