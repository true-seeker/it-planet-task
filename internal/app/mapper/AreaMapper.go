package mapper

import (
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/model/response"
)

func AreaToAreaResponse(area *entity.Area) *response.Area {
	r := response.Area{
		Id:         area.Id,
		Name:       area.Name,
		AreaPoints: *AreaPointsToAreaPointResponses(&area.AreaPoints),
	}

	return &r
}

func AreasToAreaResponses(areas *[]entity.Area) *[]response.Area {
	rs := make([]response.Area, 0)

	for _, area := range *areas {
		rs = append(rs, *AreaToAreaResponse(&area))
	}

	return &rs
}

func AreaResponseToArea(areaResponse *response.Area) *entity.Area {
	r := entity.Area{
		Id:         areaResponse.Id,
		Name:       areaResponse.Name,
		AreaPoints: *AreaPointResponsesToAreaPoints(&areaResponse.AreaPoints),
	}

	return &r
}
