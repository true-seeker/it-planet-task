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

func DecodeCredentials(c *gin.Context) (string, string, bool) {
	r := c.Request
	return r.BasicAuth()
}

func GetAccountByCreds(c *gin.Context) (*entity.Account, error) {
	login, password, ok := DecodeCredentials(c)
	if !ok {
		return nil, errors.New("")
	}

	account := &entity.Account{
		Email:    login,
		Password: password,
	}

	accountRepo := repository.NewAccountRepository(helpers.GetConnectionOrCreateAndGet())
	accountService := service.NewAccountService(accountRepo)

	return accountService.GetByCreds(account), nil

}

// TokenAuth middleware для basic auth
func TokenAuth(c *gin.Context) {
	token := c.Request.Header.Get("token")
	authRepo := repository.NewAuthRepository(helpers.GetConnectionOrCreateAndGet())
	accountService := service.NewAuthService(authRepo)
	isAuthed, _ := accountService.CheckToken(token)
	if !isAuthed {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
