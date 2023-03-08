package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/input"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AnimalValidator"
	"net/http"
)

type AnimalHandler struct {
	service               service.Animal
	animalTypeService     service.AnimalType
	accountService        service.Account
	locationService       service.Location
	animalLocationService service.AnimalLocation
}

func NewAnimalHandler(service service.Animal, animalTypeService service.AnimalType, accountService service.Account, locationService service.Location, animalLocationService service.AnimalLocation) *AnimalHandler {
	return &AnimalHandler{service: service, animalTypeService: animalTypeService, accountService: accountService, locationService: locationService, animalLocationService: animalLocationService}
}

func (a *AnimalHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animal, httpErr := a.service.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
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

	_, httpErr = a.accountService.Get(newAnimal.ChipperId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	_, httpErr = a.locationService.Get(newAnimal.ChippingLocationId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
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

	oldAnimal, httpErr := a.service.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	httpErr = AnimalValidator.ValidateAnimalUpdateInput(animalInput, oldAnimal)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	if len(oldAnimal.VisitedLocationsId) > 0 {
		firstVisitedLocation, httpErr := a.animalLocationService.Get(oldAnimal.VisitedLocationsId[0])
		if httpErr != nil {
			c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
			return
		}
		if *animalInput.ChippingLocationId == firstVisitedLocation.LocationPointId {
			c.AbortWithStatusJSON(http.StatusBadRequest, "new chipping location matches with first visited point")
			return
		}
	}

	animalInput.AnimalTypeIds = oldAnimal.AnimalTypesId
	newAnimal := mapper.AnimalInputToAnimal(animalInput)

	_, httpErr = a.accountService.Get(newAnimal.ChipperId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	_, httpErr = a.locationService.Get(newAnimal.ChippingLocationId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
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

	animalResponse, httpErr := a.service.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	if len(animalResponse.VisitedLocationsId) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "animal has visited location points")
		return
	}

	err := a.service.Delete(id)
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

	animalResponse, httpErr := a.service.Get(animalId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	_, httpErr = a.animalTypeService.Get(typeId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	for _, elem := range animalResponse.AnimalTypesId {
		if elem == typeId {
			c.AbortWithStatusJSON(http.StatusConflict, fmt.Sprintf("Animal already has type with id %d", typeId))
			return
		}
	}

	animalResponse, err := a.service.AddAnimalType(animalId, typeId)
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

	animalResponse, httpErr := a.service.Get(animalId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	_, httpErr = a.animalTypeService.Get(*animalTypeUpdateInput.NewTypeId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	_, httpErr = a.animalTypeService.Get(*animalTypeUpdateInput.OldTypeId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
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

	animalResponse, httpErr := a.service.Get(animalId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	_, httpErr = a.animalTypeService.Get(typeId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
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

	animalResponse, err := a.service.DeleteAnimalType(animalId, typeId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, animalResponse)
}
