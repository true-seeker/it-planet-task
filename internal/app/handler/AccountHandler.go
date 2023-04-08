package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AccountValidator"
	"net/http"
)

// AccountHandler Обработчик запросов для сущности "Аккаунт"
type AccountHandler struct {
	accountService service.Account
	animalService  service.Animal
}

func NewAccountHandler(accountService service.Account, animalService service.Animal) *AccountHandler {
	return &AccountHandler{accountService: accountService, animalService: animalService}
}

func (a *AccountHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	authorizedAccountAny, _ := c.Get("account")
	authorizedAccount := authorizedAccountAny.(*entity.Account)
	if (authorizedAccount.Role == entity.UserRole || authorizedAccount.Role == entity.ChipperRole) && (id != authorizedAccount.Id) {
		c.AbortWithStatusJSON(http.StatusForbidden, "Cant get another's account")
		return
	}

	account, httpErr := a.accountService.Get(id)
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

	accounts, err := a.accountService.Search(params)
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
	authorizedAccountAny, _ := c.Get("account")
	authorizedAccount := authorizedAccountAny.(*entity.Account)
	if (authorizedAccount.Role == entity.UserRole || authorizedAccount.Role == entity.ChipperRole) && (id != authorizedAccount.Id) {
		c.AbortWithStatusJSON(http.StatusForbidden, "Cant edit another's account")
		return
	}

	newAccount := &entity.Account{}
	err := c.BindJSON(&newAccount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	httpErr = AccountValidator.ValidateAccount(newAccount)
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}

	duplicateAccount, err := a.accountService.GetByEmail(newAccount)
	if duplicateAccount.Id != 0 {
		if duplicateAccount.Id != id {
			c.AbortWithStatusJSON(http.StatusConflict, fmt.Sprintf("Account with email %s already exists", newAccount.Email))
			return
		}
	}

	newAccount.Id = id
	account, err := a.accountService.Update(newAccount)
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
	authorizedAccountAny, _ := c.Get("account")
	authorizedAccount := authorizedAccountAny.(*entity.Account)
	if (authorizedAccount.Role == entity.UserRole || authorizedAccount.Role == entity.ChipperRole) && (id != authorizedAccount.Id) {
		c.AbortWithStatusJSON(http.StatusForbidden, "Cant delete another's account")
		return
	}

	animals, _ := a.animalService.GetAnimalsByAccountId(id)
	if len(*animals) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Account has animals attached")
		return
	}

	err := a.accountService.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (a *AccountHandler) Create(c *gin.Context) {
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

	duplicateAccount, err := a.accountService.GetByEmail(newAccount)
	if duplicateAccount.Id != 0 {
		c.AbortWithStatusJSON(http.StatusConflict, fmt.Sprintf("Account with email %s already exists", newAccount.Email))
		return
	}

	account, err := a.accountService.Create(newAccount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, account)
}
