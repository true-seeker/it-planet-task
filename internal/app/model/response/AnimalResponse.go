package response

import "time"

type Animal struct {
	Id                 int        `json:"id"`
	AnimalTypesId      []int      `json:"animalTypes"`
	Weight             float32    `json:"weight"`
	Height             float32    `json:"height"`
	Length             float32    `json:"length"`
	Gender             string     `json:"gender"`
	LifeStatus         string     `json:"lifeStatus"`
	ChippingDateTime   time.Time  `json:"chippingDateTime"`
	ChipperId          int        `json:"chipperId"`
	ChippingLocationId int        `json:"chippingLocationId"`
	VisitedLocationsId []int      `json:"visitedLocations"`
	DeathDateTime      *time.Time `json:"deathDateTime"`
}

type AnimalForAreaAnalyticsDTO struct {
	AnimalId     int    `json:"animal_id"`
	Type         string `json:"type"`
	AnimalTypeId int    `json:"animal_type_id"`
}
