package handler

import (
	"github.com/gin-gonic/gin"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator/AccountValidator"
	"it-planet-task/internal/pkg/middleware"
	"net/http"
)

type AuthHandler struct {
	service        service.Auth
	accountService service.Account
}

func NewAuthHandler(service service.Auth, accountService service.Account) *AuthHandler {
	return &AuthHandler{service: service, accountService: accountService}
}

func (a *AuthHandler) Register(c *gin.Context) {
	isAuthenticated, _ := middleware.IsAccountExists(c)
	if isAuthenticated {
		c.AbortWithStatusJSON(http.StatusForbidden, "Already authenticated")
		return
	}

	newAccount := &entity.Account{}
	err := c.BindJSON(&newAccount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	httpErr := AccountValidator.ValidateAccount(newAccount)
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}

	if a.accountService.IsAlreadyExists(newAccount) {
		c.AbortWithStatusJSON(http.StatusConflict, "Account already exists")
		return
	}

	account, err := a.service.Register(newAccount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, account)
}
