package AreaValidator

import (
	"fmt"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service/geometry"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AreaPointValidator"
	"it-planet-task/pkg/errorHandler"
	"net/http"
)

func ValidateArea(area *entity.Area) *errorHandler.HttpErr {
	// TODO совместить сообщения об ошибках
	if validator.IsStringEmpty(area.Name) {
		return errorHandler.NewHttpErr("name is empty", http.StatusBadRequest)
	}

	if area.AreaPoints == nil {
		return errorHandler.NewHttpErr("areaPoints is empty", http.StatusBadRequest)
	}

	if len(area.AreaPoints) < 3 {
		return errorHandler.NewHttpErr("size of areaPoints must be greater than 2", http.StatusBadRequest)
	}
	for _, areaPoint := range area.AreaPoints {
		httpErr := AreaPointValidator.ValidateAreaPoint(&areaPoint)
		if httpErr != nil {
			return httpErr
		}
	}
	geometryService := geometry.NewGeometryService()

	if geometryService.IsAllPointsOnOneLine(&area.AreaPoints) {
		return errorHandler.NewHttpErr("all points are on the same line", http.StatusBadRequest)
	}

	for i := 0; i < len(area.AreaPoints); i++ {
		for j := i + 1; j < len(area.AreaPoints); j++ {
			if area.AreaPoints[i].IsEqual(&area.AreaPoints[j]) {
				return errorHandler.NewHttpErr("area contains duplicate points", http.StatusBadRequest)
			}
		}
	}

	lineSegments := make([]geometry.LineSegment, 0)
	for i := 0; i < len(area.AreaPoints)-1; i++ {
		lineSegments = append(lineSegments, *geometry.NewLineSegment(area.AreaPoints[i], area.AreaPoints[i+1]))
	}
	lineSegments = append(lineSegments, *geometry.NewLineSegment(area.AreaPoints[0], area.AreaPoints[len(area.AreaPoints)-1]))

	for i := 0; i < len(lineSegments); i++ {
		for j := i + 1; j < len(lineSegments); j++ {
			if geometryService.IsIntersects(&lineSegments[i], &lineSegments[j]) {
				return errorHandler.NewHttpErr("area must be non self-intersecting", http.StatusBadRequest)
			}
		}
	}

	return nil
}

func ValidateIntersectionAndAreaRepeats(area *entity.Area, existingArea *entity.Area) *errorHandler.HttpErr {
	geometryService := geometry.NewGeometryService()

	if area.Name == existingArea.Name {
		return errorHandler.NewHttpErr("area with this name already exists", http.StatusConflict)
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
			if i == len(existingArea.AreaPoints) {
				i = 0
			}
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
			return errorHandler.NewHttpErr(fmt.Sprintf("area with these points already exists with id %d", existingArea.Id), http.StatusConflict)
		}

	}

	areaLineSegments := make([]geometry.LineSegment, 0)
	for i := 0; i < len(area.AreaPoints)-1; i++ {
		areaLineSegments = append(areaLineSegments, *geometry.NewLineSegment(area.AreaPoints[i], area.AreaPoints[i+1]))
	}
	if len(areaLineSegments) > 0 {
		areaLineSegments = append(areaLineSegments, *geometry.NewLineSegment(area.AreaPoints[0], area.AreaPoints[len(area.AreaPoints)-1]))
	}

	existingLineSegments := make([]geometry.LineSegment, 0)
	for i := 0; i < len(existingArea.AreaPoints)-1; i++ {
		existingLineSegments = append(existingLineSegments, *geometry.NewLineSegment(existingArea.AreaPoints[i], existingArea.AreaPoints[i+1]))
	}
	if len(existingLineSegments) > 0 {
		existingLineSegments = append(existingLineSegments, *geometry.NewLineSegment(existingArea.AreaPoints[0], existingArea.AreaPoints[len(existingArea.AreaPoints)-1]))
	}

	for i := 0; i < len(existingLineSegments); i++ {
		for j := 0; j < len(areaLineSegments); j++ {
			if geometryService.IsIntersects(&existingLineSegments[i], &areaLineSegments[j]) {
				return errorHandler.NewHttpErr(fmt.Sprintf("area intersects with area with id %d", existingArea.Id), http.StatusBadRequest)
			}
		}
	}

	for _, point := range area.AreaPoints {
		if geometryService.IsPointInsideArea(&point, existingArea, false) {
			return errorHandler.NewHttpErr(fmt.Sprintf("area lays inside area with id %d", existingArea.Id), http.StatusBadRequest)
		}
	}

	for _, point := range existingArea.AreaPoints {
		if geometryService.IsPointInsideArea(&point, area, false) {
			return errorHandler.NewHttpErr(fmt.Sprintf("area contains area with id %d", existingArea.Id), http.StatusBadRequest)
		}
	}
	return nil
}
