package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mmcloughlin/geohash"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/LocationValidator"
	"net/http"
)

// LocationHandler Обработчик запросов для сущности "Локация"
type LocationHandler struct {
	locationService service.Location
	animalService   service.Animal
}

func NewLocationHandler(locationService service.Location, animalService service.Animal) *LocationHandler {
	return &LocationHandler{locationService: locationService, animalService: animalService}
}

func (l *LocationHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	location, httpErr := l.locationService.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	c.JSON(http.StatusOK, location)
}

func (l *LocationHandler) GetByCoordinates(c *gin.Context) {
	params, httpErr := filter.NewLocationCoordinatesParams(c.Request.URL.Query())
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	location := &entity.Location{
		Latitude:  params.Latitude,
		Longitude: params.Longitude,
	}

	httpErr = LocationValidator.ValidateLocation(location)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	locationResponse, httpErr := l.locationService.GetByCoordinates(location)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	c.JSON(http.StatusOK, locationResponse.Id)
}

func (l *LocationHandler) GeoHash(c *gin.Context) {
	params, httpErr := filter.NewLocationCoordinatesParams(c.Request.URL.Query())
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	location := &entity.Location{
		Latitude:  params.Latitude,
		Longitude: params.Longitude,
	}

	httpErr = LocationValidator.ValidateLocation(location)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	// todo geohash db
	a := geohash.Encode(*params.Latitude, *params.Longitude)
	c.String(http.StatusOK, a)
}

func (l *LocationHandler) GeoHashV2(c *gin.Context) {
	params, httpErr := filter.NewLocationCoordinatesParams(c.Request.URL.Query())
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	location := &entity.Location{
		Latitude:  params.Latitude,
		Longitude: params.Longitude,
	}

	httpErr = LocationValidator.ValidateLocation(location)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	c.JSON(http.StatusOK, geohash.Encode(*params.Latitude, *params.Longitude))
}

func (l *LocationHandler) GeoHashV3(c *gin.Context) {
	params, httpErr := filter.NewLocationCoordinatesParams(c.Request.URL.Query())
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	location := &entity.Location{
		Latitude:  params.Latitude,
		Longitude: params.Longitude,
	}

	httpErr = LocationValidator.ValidateLocation(location)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	c.JSON(http.StatusOK, geohash.Encode(*params.Latitude, *params.Longitude))
}

func (l *LocationHandler) Create(c *gin.Context) {
	newLocation := &entity.Location{}
	err := c.BindJSON(&newLocation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	httpErr := LocationValidator.ValidateLocation(newLocation)
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}

	_, httpErr = l.locationService.GetByCoordinates(newLocation)
	if httpErr == nil {
		c.AbortWithStatusJSON(http.StatusConflict, "Location already exists")
		return
	}

	location, err := l.locationService.Create(newLocation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, location)
}

func (l *LocationHandler) Update(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	oldLocation, httpErr := l.locationService.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	newLocation := &entity.Location{}
	err := c.BindJSON(&newLocation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	httpErr = LocationValidator.ValidateLocation(newLocation)
	if httpErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, httpErr.Err.Error())
		return
	}

	duplicateLocation, httpErr := l.locationService.GetByCoordinates(newLocation)
	if httpErr == nil && oldLocation.Id != duplicateLocation.Id {
		c.AbortWithStatusJSON(http.StatusConflict, "Location already exists")
		return
	}

	newLocation.Id = id
	location, err := l.locationService.Update(newLocation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, location)
}

func (l *LocationHandler) Delete(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	_, httpErr = l.locationService.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	animals, err := l.animalService.GetAnimalsByLocationId(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if len(*animals) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("There are animals with loocation id %d", id))
		return
	}

	err = l.locationService.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
