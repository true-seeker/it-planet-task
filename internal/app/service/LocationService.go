package service

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"it-planet-task/internal/app/mapper"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
	"it-planet-task/internal/app/repository"
	"it-planet-task/internal/app/service/geohash"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

type Location interface {
	Get(id int) (*response.Location, *errorHandler.HttpErr)
	Create(location *entity.Location) (*response.Location, error)
	Update(location *entity.Location) (*response.Location, error)
	Delete(id int) error
	GetByCoordinates(location *entity.Location) (*response.Location, *errorHandler.HttpErr)
	GeoHashV1(location *entity.Location) (*string, *errorHandler.HttpErr)
	GeoHashV2(location *entity.Location) (*string, *errorHandler.HttpErr)
	GeoHashV3(location *entity.Location) (*string, *errorHandler.HttpErr)
}

type LocationService struct {
	locationRepo repository.Location
}

func NewLocationService(serviceRepo repository.Location) Location {
	return &LocationService{locationRepo: serviceRepo}
}

func (l *LocationService) Get(id int) (*response.Location, *errorHandler.HttpErr) {
	locationResponse := &response.Location{}

	location, err := l.locationRepo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorHandler.NewHttpErr(fmt.Sprintf("Location with id %d does not exists", id), http.StatusNotFound)
		} else {
			return nil, errorHandler.NewHttpErr(err.Error(), http.StatusBadRequest)
		}
	}

	locationResponse = mapper.LocationToLocationResponse(location)

	return locationResponse, nil
}

func (l *LocationService) Create(location *entity.Location) (*response.Location, error) {
	locationResponse := &response.Location{}

	location, err := l.locationRepo.Create(location)
	if err != nil {
		return nil, err
	}

	locationResponse = mapper.LocationToLocationResponse(location)

	return locationResponse, nil
}

func (l *LocationService) Update(location *entity.Location) (*response.Location, error) {
	locationResponse := &response.Location{}

	location, err := l.locationRepo.Update(location)
	if err != nil {
		return nil, err
	}

	locationResponse = mapper.LocationToLocationResponse(location)

	return locationResponse, nil
}

func (l *LocationService) Delete(id int) error {
	return l.locationRepo.Delete(id)
}

func (l *LocationService) GetByCoordinates(location *entity.Location) (*response.Location, *errorHandler.HttpErr) {
	locationResponse := &response.Location{}
	location, err := l.locationRepo.GetByCoordinates(location)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorHandler.NewHttpErr("location with these coordinates does not exists", http.StatusNotFound)
		} else {
			return nil, errorHandler.NewHttpErr(err.Error(), http.StatusBadRequest)
		}
	}

	locationResponse = mapper.LocationToLocationResponse(location)

	return locationResponse, nil
}

func (l *LocationService) GeoHashV1(location *entity.Location) (*string, *errorHandler.HttpErr) {
	geoHashV1 := geohash.Encode(*location.Latitude, *location.Longitude)

	return &geoHashV1, nil

}

func (l *LocationService) GeoHashV2(location *entity.Location) (*string, *errorHandler.HttpErr) {
	geoHashV1, httpErr := l.GeoHashV1(location)
	if httpErr != nil {
		return nil, httpErr
	}

	geoHashV2 := base64.StdEncoding.EncodeToString([]byte(*geoHashV1))

	return &geoHashV2, nil
}

func (l *LocationService) GeoHashV3(location *entity.Location) (*string, *errorHandler.HttpErr) {
	geoHashV1, httpErr := l.GeoHashV1(location)
	if httpErr != nil {
		return nil, httpErr
	}

	md := md5.Sum([]byte(*geoHashV1))

	reversedMd5 := make([]byte, 16)
	for i := 0; i < len(md); i++ {
		reversedMd5[len(md)-i-1] = md[i]
	}

	geoHashV3 := base64.StdEncoding.EncodeToString(reversedMd5)

	return &geoHashV3, nil
}
