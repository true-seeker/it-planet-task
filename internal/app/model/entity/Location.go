package entity

type Location struct {
	Id        int      `gorm:"primary_key"`
	Latitude  *float64 `gorm:"not_null"`
	Longitude *float64 `gorm:"not_null"`
}

func NewLocation(id int, latitude *float64, longitude *float64) *Location {
	return &Location{Id: id, Latitude: latitude, Longitude: longitude}
}
