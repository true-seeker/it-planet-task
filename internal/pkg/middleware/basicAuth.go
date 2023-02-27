package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"it-planet-task/helpers"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/repository"
	"it-planet-task/internal/app/service"
	"net/http"
)

func IsAccountExists(c *gin.Context) (bool, error) {
	r := c.Request
	login, password, ok := r.BasicAuth()
	if !ok {
		return false, errors.New("")
	}

	account := &entity.Account{
		Email:    login,
		Password: password,
	}

	accountRepo := repository.NewAccountRepository(helpers.GetConnectionOrCreateAndGet())
	accountService := service.NewAccountService(accountRepo)

	return accountService.CheckCredentials(account), nil

}

func BasicAuth(c *gin.Context) {
	isExists, err := IsAccountExists(c)
	if err != nil || !isExists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Next()
}
