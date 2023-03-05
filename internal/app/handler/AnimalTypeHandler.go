package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AnimalTypeValidator"
	"net/http"
)

type AnimalTypeHandler struct {
	service       service.AnimalType
	animalService service.Animal
}

func NewAnimalTypeHandler(service service.AnimalType, animalService service.Animal) *AnimalTypeHandler {
	return &AnimalTypeHandler{service: service, animalService: animalService}
}

func (a *AnimalTypeHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "id must be greater than 0")
		return
	}

	animalType, err := a.service.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	}
	c.JSON(http.StatusOK, animalType)
}

func (a *AnimalTypeHandler) Create(c *gin.Context) {
	newAnimalType := &entity.AnimalType{}
	err := c.BindJSON(&newAnimalType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	httpErr := AnimalTypeValidator.ValidateAnimalType(newAnimalType)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	duplicateAnimalType := a.service.GetByType(newAnimalType)
	if duplicateAnimalType.Id != 0 {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	animalType, err := a.service.Create(newAnimalType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, animalType)
}

func (a *AnimalTypeHandler) Update(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "id must be greater than 0")
		return
	}

	newAnimalType := &entity.AnimalType{}
	err := c.BindJSON(&newAnimalType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	duplicateAnimalType := a.service.GetByType(newAnimalType)
	if duplicateAnimalType.Id != 0 && id != duplicateAnimalType.Id {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	_, err = a.service.Get(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	httpErr = AnimalTypeValidator.ValidateAnimalType(newAnimalType)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err)
		return
	}

	newAnimalType.Id = id
	animalType, _ := a.service.Update(newAnimalType)

	c.JSON(http.StatusOK, animalType)

}

func (a *AnimalTypeHandler) Delete(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "id must be greater than 0")
		return
	}

	_, err := a.service.Get(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	animals, _ := a.animalService.GetAnimalsByAnimalTypeId(id)
	if len(*animals) != 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = a.service.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusOK)
}
