package entity

import "time"

type AnimalLocation struct {
	Id                           int `gorm:"primary_key"`
	DateTimeOfVisitLocationPoint time.Time
	LocationPointId              int
	LocationPoint                Location
}
