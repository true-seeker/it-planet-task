package mapper

import (
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
)

func LocationToLocationResponse(location *entity.Location) *response.Location {
	r := response.Location{
		Id:        location.Id,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	}

	return &r
}

func LocationsToLocationResponses(locations *[]entity.Location) *[]response.Location {
	rs := make([]response.Location, 0)

	for _, location := range *locations {
		rs = append(rs, *LocationToLocationResponse(&location))
	}

	return &rs
}
