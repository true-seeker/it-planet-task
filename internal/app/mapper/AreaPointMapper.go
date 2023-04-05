package mapper

import (
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
)

func AreaPointToAreaPointResponse(areaPoint *entity.AreaPoint) *response.AreaPoint {
	r := response.AreaPoint{
		Latitude:  areaPoint.Latitude,
		Longitude: areaPoint.Longitude,
	}
	return &r
}

func AreaPointsToAreaPointResponses(areaPoints *[]entity.AreaPoint) *[]response.AreaPoint {
	rs := make([]response.AreaPoint, 0)

	for _, areaPoint := range *areaPoints {
		rs = append(rs, *AreaPointToAreaPointResponse(&areaPoint))
	}

	return &rs
}

func AreaPointResponseToAreaPoint(areaPointResponse *response.AreaPoint) *entity.AreaPoint {
	r := entity.AreaPoint{
		Latitude:  areaPointResponse.Latitude,
		Longitude: areaPointResponse.Longitude,
	}
	return &r
}

func AreaPointResponsesToAreaPoints(areaPointResponse *[]response.AreaPoint) *[]entity.AreaPoint {
	rs := make([]entity.AreaPoint, 0)

	for _, areaPointResponse := range *areaPointResponse {
		rs = append(rs, *AreaPointResponseToAreaPoint(&areaPointResponse))
	}

	return &rs
}

func LocationToAreaPoint(location *entity.Location) *entity.AreaPoint {
	areaPoint := entity.AreaPoint{
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	}
	return &areaPoint
}
