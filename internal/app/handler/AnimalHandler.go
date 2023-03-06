package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/input"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AnimalValidator"
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
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animal, err := a.service.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Animal with id %d does not exists", id))
			return
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
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
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, animals)
}

func (a *AnimalHandler) Create(c *gin.Context) {
	animalInput := &input.Animal{}
	err := c.BindJSON(&animalInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
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
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Chipper with id %d does not exists", newAnimal.ChipperId))
		return
	}

	_, err = a.locationService.Get(newAnimal.ChippingLocationId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Location with id %d does not exists", newAnimal.ChippingLocationId))
		return
	}

	animalTypes, err := a.animalTypeService.GetByIds(&animalInput.AnimalTypeIds)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if len(*animalTypes) < len(animalInput.AnimalTypeIds) {
		c.AbortWithStatusJSON(http.StatusNotFound, "Some animal types does not exist")
		return
	}

	animal, err := a.service.Create(newAnimal)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, animal)
}

func (a *AnimalHandler) Update(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalInput := &input.Animal{}
	err := c.BindJSON(&animalInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	oldAnimal, err := a.service.Get(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Animal with id %d does not exists", id))
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
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Chipper with id %d does not exists", newAnimal.ChipperId))
		return
	}

	_, err = a.locationService.Get(newAnimal.ChippingLocationId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Location with id %d does not exists", newAnimal.ChippingLocationId))
		return
	}

	newAnimal.Id = oldAnimal.Id

	animalResponse, err := a.service.Update(newAnimal, oldAnimal)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, animalResponse)
}

func (a *AnimalHandler) Delete(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	// todo Животное покинуло локацию чипирования, при этом есть другие посещенные точки

	_, err := a.service.Get(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Animal with id %d does not exists", id))
		return
	}

	err = a.service.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (a *AnimalHandler) AddAnimalType(c *gin.Context) {
	animalId, httpErr := validator.ValidateAndReturnId(c.Param("id"), "animalId")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	typeId, httpErr := validator.ValidateAndReturnId(c.Param("typeId"), "typeId")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalResponse, err := a.service.Get(animalId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Animal with id %d does not exists", animalId))
		return
	}

	_, err = a.animalTypeService.Get(typeId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Animal type with id %d does not exists", typeId))
		return
	}

	for _, elem := range animalResponse.AnimalTypesId {
		if elem == typeId {
			c.AbortWithStatusJSON(http.StatusConflict, fmt.Sprintf("Animal already has type with id %d", typeId))
			return
		}
	}

	animalResponse, err = a.service.AddAnimalType(animalId, typeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, animalResponse)
}

func (a *AnimalHandler) EditAnimalType(c *gin.Context) {
	animalId, httpErr := validator.ValidateAndReturnId(c.Param("id"), "animalId")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalTypeUpdateInput := &input.AnimalTypeUpdate{}
	err := c.BindJSON(&animalTypeUpdateInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	httpErr = AnimalValidator.ValidateAnimalTypeUpdateInput(animalTypeUpdateInput)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalResponse, err := a.service.Get(animalId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Animal with id %d does not exists", animalId))
		return
	}

	_, err = a.animalTypeService.Get(*animalTypeUpdateInput.NewTypeId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Animal type with id %d does not exists", *animalTypeUpdateInput.NewTypeId))
		return
	}
	_, err = a.animalTypeService.Get(*animalTypeUpdateInput.OldTypeId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Animal type with id %d does not exists", *animalTypeUpdateInput.OldTypeId))
		return
	}

	animalTypeIdsMap := make(map[int]bool)
	for _, elem := range animalResponse.AnimalTypesId {
		if elem == *animalTypeUpdateInput.NewTypeId {
			c.AbortWithStatusJSON(http.StatusConflict, fmt.Sprintf("Animal already has type with id %d", *animalTypeUpdateInput.NewTypeId))
			return
		}
		animalTypeIdsMap[elem] = true
	}
	if !animalTypeIdsMap[*animalTypeUpdateInput.OldTypeId] {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Animal does not have type with id %d", *animalTypeUpdateInput.OldTypeId))
		return
	}

	animalResponse, err = a.service.EditAnimalType(animalId, animalTypeUpdateInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, animalResponse)
}

func (a *AnimalHandler) DeleteAnimalType(c *gin.Context) {
	animalId, httpErr := validator.ValidateAndReturnId(c.Param("id"), "animalId")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	typeId, httpErr := validator.ValidateAndReturnId(c.Param("typeId"), "typeId")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animalResponse, err := a.service.Get(animalId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Animal with id %d does not exists", animalId))
		return
	}

	_, err = a.animalTypeService.Get(typeId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Animal type with id %d does not exists", typeId))
		return
	}

	animalTypeIdsMap := make(map[int]bool)
	for _, elem := range animalResponse.AnimalTypesId {
		animalTypeIdsMap[elem] = true
	}
	if !animalTypeIdsMap[typeId] {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Animal does not have type with id %d", typeId))
		return
	}

	if len(animalResponse.AnimalTypesId) == 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Animal has only 1 type")
		return
	}

	animalResponse, err = a.service.DeleteAnimalType(animalId, typeId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, animalResponse)
}
