package handler

import (
	"github.com/gin-gonic/gin"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator/AccountValidator"
	"net/http"
)

// AuthHandler Обработчик запросов, связанных с аутентификацией
type AuthHandler struct {
	authService    service.Auth
	accountService service.Account
}

func NewAuthHandler(authService service.Auth, accountService service.Account) *AuthHandler {
	return &AuthHandler{authService: authService, accountService: accountService}
}

func (a *AuthHandler) Register(c *gin.Context) {
	authenticatedAccount, _ := c.Get("account")
	if authenticatedAccount != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, "Already authenticated")
		return
	}

	newAccount := &entity.Account{}
	err := c.BindJSON(&newAccount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	httpErr := AccountValidator.ValidateAccountRegistration(newAccount)
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}

	if a.accountService.IsAlreadyExists(newAccount) {
		c.AbortWithStatusJSON(http.StatusConflict, "Account already exists")
		return
	}

	account, err := a.authService.Register(newAccount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, account)
}

func (a *AuthHandler) Login(c *gin.Context) {
	account := &entity.Account{}
	err := c.BindJSON(&account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	account = a.accountService.GetByCreds(account)
	if account.Id == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken, httpErr := a.authService.Login(account)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	c.JSON(http.StatusOK, authToken)
}
