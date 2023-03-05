package response

import "time"

type AnimalLocation struct {
	Id                           int       `json:"id"`
	DateTimeOfVisitLocationPoint time.Time `json:"dateTimeOfVisitLocationPoint"`
	LocationPointId              int       `json:"locationPointId"`
}
