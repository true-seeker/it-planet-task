package response

import "time"

type AnimalLocation struct {
	Id                           int       `json:"id"`
	DateTimeOfVisitLocationPoint time.Time `json:"dateTimeOfVisitLocationPoint"`
	LocationPointId              int       `json:"locationPointId"`
}

type AnimalLocationForAreaAnalyticsDTO struct {
	AnimalId                     int       `json:"animal_id"`
	DateTimeOfVisitLocationPoint time.Time `json:"dateTimeOfVisitLocationPoint"`
	Latitude                     *float64  `json:"latitude"`
	Longitude                    *float64  `json:"longitude"`
	IsPrevious                   bool      `json:"isPrevious"`
}
