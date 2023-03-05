package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/input"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AnimalValidator"
	"it-planet-task/pkg/converter"
	"net/http"
)

type AnimalHandler struct {
	service           service.Animal
	animalTypeService service.AnimalType
	accountService    service.Account
	locationService   service.Location
}

func NewAnimalHandler(service service.Animal, animalTypeService service.AnimalType, accountService service.Account, locationService service.Location) *AnimalHandler {
	return &AnimalHandler{service: service, animalTypeService: animalTypeService, accountService: accountService, locationService: locationService}
}

func (a *AnimalHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "id must be greater than 0")
		return
	}

	animal, err := a.service.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	}
	c.JSON(http.StatusOK, animal)
}

func (a *AnimalHandler) GetAnimalLocations(c *gin.Context) {
	id, err := converter.StringToInt(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	animal, err := a.service.GetAnimalLocations(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return

		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	}
	c.JSON(http.StatusOK, animal)
}

func (a *AnimalHandler) Search(c *gin.Context) {
	params, httpErr := filter.NewAnimalFilterParams(c.Request.URL.Query())
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	animals, err := a.service.Search(params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, animals)
}

func (a *AnimalHandler) Create(c *gin.Context) {
	animalInput := &input.Animal{}
	err := c.BindJSON(&animalInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	httpErr := AnimalValidator.ValidateAnimalCreateInput(animalInput)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	newAnimal := mapper.AnimalInputToAnimal(animalInput)

	_, err = a.accountService.Get(newAnimal.ChipperId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	_, err = a.locationService.Get(newAnimal.ChippingLocationId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	animalTypes, err := a.animalTypeService.GetByIds(&animalInput.AnimalTypeIds)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	if len(*animalTypes) < len(animalInput.AnimalTypeIds) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	animal, err := a.service.Create(newAnimal)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, animal)
}

func (a *AnimalHandler) Update(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "id must be greater than 0")
		return
	}

	animalInput := &input.Animal{}
	err := c.BindJSON(&animalInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	oldAnimal, err := a.service.Get(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	httpErr = AnimalValidator.ValidateAnimalUpdateInput(animalInput, oldAnimal)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalInput.AnimalTypeIds = oldAnimal.AnimalTypesId
	newAnimal := mapper.AnimalInputToAnimal(animalInput)

	_, err = a.accountService.Get(newAnimal.ChipperId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	_, err = a.locationService.Get(newAnimal.ChippingLocationId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	newAnimal.Id = oldAnimal.Id

	animalResponse, err := a.service.Update(newAnimal, oldAnimal)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, animalResponse)
}

func (a *AnimalHandler) Delete(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnIntField(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "id must be greater than 0")
		return
	}
	// todo Животное покинуло локацию чипирования, при этом есть другие посещенные точки

	_, err := a.service.Get(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	err = a.service.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusOK)
}
