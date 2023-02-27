package middleware

import (
	"github.com/gin-gonic/gin"
	"it-planet-task/helpers"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/repository"
	"it-planet-task/internal/app/service"
	"net/http"
)

func BasicAuth(c *gin.Context) {
	r := c.Request
	login, password, ok := r.BasicAuth()
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	account := &entity.Account{
		Email:    login,
		Password: password,
	}

	accountRepo := repository.NewAccountRepository(helpers.GetConnectionOrCreateAndGet())
	accountService := service.NewAccountService(accountRepo)

	isExists := accountService.CheckCredentials(account)

	if !isExists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Next()
}
