package entity

type Location struct {
	Id        int      `gorm:"primary_key"`
	Latitude  *float64 `gorm:"not_null"`
	Longitude *float64 `gorm:"not_null"`
}
