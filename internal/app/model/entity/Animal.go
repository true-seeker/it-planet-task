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
	// TODO убрать линковочную таблицу
	VisitedLocations []AnimalLocation `gorm:"many2many:animal_visited_locations;not_null;constraint:OnDelete:CASCADE;"`
	DeathDateTime    *time.Time
}
