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

	if service.IsAllPointsOnOneLine(&area.AreaPoints) {
		return &errorHandler.HttpErr{
			Err:        errors.New("all points are on the same line"),
			StatusCode: http.StatusBadRequest,
		}
	}

	for i := 0; i < len(area.AreaPoints); i++ {
		for j := i + 1; j < len(area.AreaPoints); j++ {
			if area.AreaPoints[i].IsEqual(&area.AreaPoints[j]) {
				return &errorHandler.HttpErr{
					Err:        errors.New("area contains duplicate points"),
					StatusCode: http.StatusBadRequest,
				}
			}
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

func ValidateIntersectionAndAreaRepeats(area *entity.Area, existingArea *entity.Area) *errorHandler.HttpErr {
	if area.Name == existingArea.Name {
		return &errorHandler.HttpErr{
			Err:        errors.New("area with this name already exists"),
			StatusCode: http.StatusConflict,
		}
	}

	if len(area.AreaPoints) == len(existingArea.AreaPoints) {
		indexOfFirstFoundOverlap := -1
		for i := 0; i < len(existingArea.AreaPoints) && indexOfFirstFoundOverlap == -1; i++ {
			if existingArea.AreaPoints[i].IsEqual(&area.AreaPoints[0]) {
				indexOfFirstFoundOverlap = i
			}
		}

		isFullOverlapFound := true
		j := 1
		for i := indexOfFirstFoundOverlap + 1; i != indexOfFirstFoundOverlap && isFullOverlapFound; {
			if existingArea.AreaPoints[i].IsEqual(&area.AreaPoints[j]) {
				if i < len(existingArea.AreaPoints)-1 {
					i++
				} else {
					i = 0
				}
				j++
			} else {
				isFullOverlapFound = false
			}
		}

		if isFullOverlapFound {
			return &errorHandler.HttpErr{
				Err:        errors.New(fmt.Sprintf("area with these points already exists with id %d", existingArea.Id)),
				StatusCode: http.StatusConflict,
			}
		}

	}

	areaLineSegments := make([]service.LineSegment, 0)
	for i := 0; i < len(area.AreaPoints)-1; i++ {
		areaLineSegments = append(areaLineSegments, *service.NewLineSegment(area.AreaPoints[i], area.AreaPoints[i+1]))
	}
	areaLineSegments = append(areaLineSegments, *service.NewLineSegment(area.AreaPoints[0], area.AreaPoints[len(area.AreaPoints)-1]))

	existingLineSegments := make([]service.LineSegment, 0)
	for i := 0; i < len(existingArea.AreaPoints)-1; i++ {
		existingLineSegments = append(existingLineSegments, *service.NewLineSegment(existingArea.AreaPoints[i], existingArea.AreaPoints[i+1]))
	}
	existingLineSegments = append(existingLineSegments, *service.NewLineSegment(existingArea.AreaPoints[0], existingArea.AreaPoints[len(existingArea.AreaPoints)-1]))

	for i := 0; i < len(existingLineSegments); i++ {
		for j := 0; j < len(areaLineSegments); j++ {
			if existingLineSegments[i].IsIntersects(&areaLineSegments[j]) {
				return &errorHandler.HttpErr{
					Err:        errors.New(fmt.Sprintf("area intersects with area with id %d", existingArea.Id)),
					StatusCode: http.StatusBadRequest,
				}
			}
		}
	}

	for _, point := range area.AreaPoints {
		if service.IsPointInsideArea(&point, existingArea) {
			return &errorHandler.HttpErr{
				Err:        errors.New(fmt.Sprintf("area lays inside area with id %d", existingArea.Id)),
				StatusCode: http.StatusBadRequest,
			}
		}
	}

	for _, point := range existingArea.AreaPoints {
		if service.IsPointInsideArea(&point, area) {
			return &errorHandler.HttpErr{
				Err:        errors.New(fmt.Sprintf("area contains area with id %d", existingArea.Id)),
				StatusCode: http.StatusBadRequest,
			}
		}
	}
	return nil
}
