package AreaValidator

import (
	"errors"
	"fmt"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AreaPointValidator"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

func ValidateArea(area *entity.Area) *errorHandler.HttpErr {
	// TODO совместить сообщения об ошибках
	if validator.IsStringEmpty(area.Name) {
		return &errorHandler.HttpErr{
			Err:        errors.New("name is empty"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if area.AreaPoints == nil {
		return &errorHandler.HttpErr{
			Err:        errors.New("areaPoints is empty"),
			StatusCode: http.StatusBadRequest,
		}
	}

	if len(area.AreaPoints) < 3 {
		return &errorHandler.HttpErr{
			Err:        errors.New("size of areaPoints must be greater than 2"),
			StatusCode: http.StatusBadRequest,
		}
	}
	for _, areaPoint := range area.AreaPoints {
		httpErr := AreaPointValidator.ValidateAreaPoint(&areaPoint)
		if httpErr != nil {
			return httpErr
		}
	}

	lineSegments := make([]service.LineSegment, 0)
	for i := 0; i < len(area.AreaPoints)-1; i++ {
		lineSegments = append(lineSegments, *service.NewLineSegment(area.AreaPoints[i], area.AreaPoints[i+1]))
	}
	lineSegments = append(lineSegments, *service.NewLineSegment(area.AreaPoints[0], area.AreaPoints[len(area.AreaPoints)-1]))

	for i := 0; i < len(lineSegments); i++ {
		for j := i + 1; j < len(lineSegments); j++ {
			if lineSegments[i].IsIntersects(&lineSegments[j]) {
				return &errorHandler.HttpErr{
					Err:        errors.New("area must be non self-intersecting"),
					StatusCode: http.StatusBadRequest,
				}
			}
		}
	}

	return nil
}

func ValidateIntersection(area *entity.Area, existingArea *entity.Area) *errorHandler.HttpErr {
	existingLineSegments := make([]service.LineSegment, 0)
	for i := 0; i < len(existingArea.AreaPoints)-1; i++ {
		existingLineSegments = append(existingLineSegments, *service.NewLineSegment(existingArea.AreaPoints[i], existingArea.AreaPoints[i+1]))
	}
	existingLineSegments = append(existingLineSegments, *service.NewLineSegment(existingArea.AreaPoints[0], existingArea.AreaPoints[len(existingArea.AreaPoints)-1]))

	for _, point := range area.AreaPoints {
		if service.IsPointInsideArea(&point, &existingLineSegments) {
			return &errorHandler.HttpErr{
				Err:        errors.New(fmt.Sprintf("area intersects with area with id %d", existingArea.Id)),
				StatusCode: http.StatusBadRequest,
			}
		}
	}

	//for i := 0; i < len(existingLineSegments); i++ {
	//	for j := i; j < len(existingLineSegments); j++ {
	//		if existingLineSegments[i].IsIntersects(&existingLineSegments[j]) {
	//			return &errorHandler.HttpErr{
	//				Err:        errors.New("area must be non self-intersecting"),
	//				StatusCode: http.StatusBadRequest,
	//			}
	//		}
	//	}
	//}

	//polygonCandidate := service.NewPolygon(&area.AreaPoints)
	//existingPolygon := service.NewPolygon(&existingArea.AreaPoints)
	//if polygonCandidate.IsIntersect(existingPolygon) {
	//	return &errorHandler.HttpErr{
	//		Err:        errors.New(fmt.Sprintf("area intersect with existing area with id %d", existingArea.Id)),
	//		StatusCode: http.StatusBadRequest,
	//	}
	//}
	//
	return nil
}
