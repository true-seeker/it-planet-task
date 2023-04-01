package mapper

import (
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
)

func AreaPointToAreaPointResponse(areaPoint *entity.AreaPoint) *response.AreaPoint {
	r := response.AreaPoint{
		Id:        areaPoint.Id,
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
