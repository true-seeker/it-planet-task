package entity

import "time"

type Animal struct {
	Id                 int          `gorm:"primary_key"`
	AnimalTypes        []AnimalType `gorm:"many2many:animal_animal_type;not_null;constraint:OnDelete:CASCADE;"`
	Weight             float32      `gorm:"not_null"`
	Height             float32      `gorm:"not_null"`
	Length             float32      `gorm:"not_null"`
	Gender             string       `gorm:"not_null"`
	LifeStatus         string
	ChippingDateTime   time.Time
	ChipperId          int `gorm:"not_null"`
	Chipper            Account
	ChippingLocationId int `gorm:"not_null"`
	ChippingLocation   Location
	VisitedLocations   []AnimalLocation
	DeathDateTime      *time.Time
}

type AnimalLocationForAreaAnalytics struct {
	DateTimeOfVisitLocationPoint time.Time `json:"dateTimeOfVisitLocationPoint"`
	Location                     Location  `json:"location"`
	Animal                       Animal    `json:"animal"`
	IsPrevious                   bool      `json:"isPrevious"`
}
