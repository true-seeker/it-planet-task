package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AccountValidator"
	"it-planet-task/internal/pkg/middleware"
	"net/http"
)

type AccountHandler struct {
	service       service.Account
	animalService service.Animal
}

func NewAccountHandler(service service.Account, animalService service.Animal) *AccountHandler {
	return &AccountHandler{service: service, animalService: animalService}
}

func (a *AccountHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	account, httpErr := a.service.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
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
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func (a *AccountHandler) Update(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	authenticatedAccountEmail, _, _ := middleware.GetCredentials(c)
	authenticatedAccount, err := a.service.GetByEmail(&entity.Account{Email: authenticatedAccountEmail})
	if authenticatedAccount.Id != id {
		c.AbortWithStatusJSON(http.StatusForbidden, "Cant edit another's account")
		return
	}

	newAccount := &entity.Account{}
	err = c.BindJSON(&newAccount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
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
			c.AbortWithStatusJSON(http.StatusConflict, fmt.Sprintf("Account with email %s already exists", newAccount.Email))
			return
		}
	}

	newAccount.Id = authenticatedAccount.Id
	account, err := a.service.Update(newAccount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, account)
}

func (a *AccountHandler) Delete(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	authenticatedAccountEmail, _, _ := middleware.GetCredentials(c)
	authenticatedAccount, _ := a.service.GetByEmail(&entity.Account{Email: authenticatedAccountEmail})
	if authenticatedAccount.Id != id {
		c.AbortWithStatusJSON(http.StatusForbidden, "Cant delete another's account")
		return
	}

	animals, _ := a.animalService.GetAnimalsByAccountId(authenticatedAccount.Id)
	if len(*animals) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Account has animals attached")
		return
	}

	err := a.service.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
