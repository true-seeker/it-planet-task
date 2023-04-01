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

func GetCredentials(c *gin.Context) (string, string, bool) {
	r := c.Request
	return r.BasicAuth()
}

func IsAccountExists(c *gin.Context) (*entity.Account, error) {
	login, password, ok := GetCredentials(c)
	if !ok {
		return nil, errors.New("")
	}

	account := &entity.Account{
		Email:    login,
		Password: password,
	}

	accountRepo := repository.NewAccountRepository(helpers.GetConnectionOrCreateAndGet())
	accountService := service.NewAccountService(accountRepo)

	return accountService.CheckCredentials(account), nil

}

// BasicAuth middleware для basic auth
func BasicAuth(c *gin.Context) {
	acc, err := IsAccountExists(c)
	if err != nil || acc.Id == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Next()
		return
	}

	c.Set("account", acc)
	c.Next()
}
