package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AccountValidator"
	"it-planet-task/internal/pkg/middleware"
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

func (a *AccountHandler) Update(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "id must be greater than 0")
		return
	}
	authenticatedAccountEmail, _, _ := middleware.GetCredentials(c)
	authenticatedAccount, err := a.service.GetByEmail(&entity.Account{Email: authenticatedAccountEmail})
	if authenticatedAccount.Id != id {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	newAccount := &entity.Account{}
	err = c.BindJSON(&newAccount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	httpErr = AccountValidator.ValidateAccount(newAccount)
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}

	duplicateAccount, err := a.service.GetByEmail(newAccount)
	if duplicateAccount.Id != 0 {
		if duplicateAccount.Id != authenticatedAccount.Id {
			c.AbortWithStatus(http.StatusConflict)
			return
		}
	}

	newAccount.Id = authenticatedAccount.Id
	account, err := a.service.Update(newAccount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, account)
}
