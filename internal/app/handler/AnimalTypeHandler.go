package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AnimalTypeValidator"
	"net/http"
)

// AnimalTypeHandler Обработчик запросов для сущности "Тип животного"
type AnimalTypeHandler struct {
	service       service.AnimalType
	animalService service.Animal
}

func NewAnimalTypeHandler(service service.AnimalType, animalService service.Animal) *AnimalTypeHandler {
	return &AnimalTypeHandler{service: service, animalService: animalService}
}

func (a *AnimalTypeHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalType, httpErr := a.service.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	c.JSON(http.StatusOK, animalType)
}

func (a *AnimalTypeHandler) Create(c *gin.Context) {
	newAnimalType := &entity.AnimalType{}
	err := c.BindJSON(&newAnimalType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	httpErr := AnimalTypeValidator.ValidateAnimalType(newAnimalType)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	duplicateAnimalType := a.service.GetByType(newAnimalType)
	if duplicateAnimalType.Id != 0 {
		c.AbortWithStatusJSON(http.StatusConflict, fmt.Sprintf("Animal type %s already exists", newAnimalType.Type))
		return
	}

	animalType, err := a.service.Create(newAnimalType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, animalType)
}

func (a *AnimalTypeHandler) Update(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	newAnimalType := &entity.AnimalType{}
	err := c.BindJSON(&newAnimalType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	duplicateAnimalType := a.service.GetByType(newAnimalType)
	if duplicateAnimalType.Id != 0 && id != duplicateAnimalType.Id {
		c.AbortWithStatusJSON(http.StatusConflict, fmt.Sprintf("Animal type %s already exists", newAnimalType.Type))
		return
	}

	_, httpErr = a.service.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
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
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	_, httpErr = a.service.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animals, _ := a.animalService.GetAnimalsByAnimalTypeId(id)
	if len(*animals) != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("There are animals with animal type id %d", id))
		return
	}
	err := a.service.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
