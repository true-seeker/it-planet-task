package response

import "time"

type Animal struct {
	Id                 int
	AnimalTypes        []int
	Weight             float32
	Height             float32
	Length             float32
	Gender             string
	LifeStatus         string
	ChippingDateTime   time.Time
	ChipperId          int
	ChippingLocationId int
	VisitedLocations   []int
	DeathDateTime      time.Time
}
