package entity

import "time"

type AnimalLocation struct {
	Id                           int       `gorm:"primary_key"`
	DateTimeOfVisitLocationPoint time.Time `gorm:"not_null"`
	LocationPointId              int       `gorm:"not_null"`
	LocationPoint                Location
}
