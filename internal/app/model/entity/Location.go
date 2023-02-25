package entity

type Location struct {
	Id        int     `gorm:"primary_key"`
	Latitude  float32 `gorm:"not_null"`
	Longitude float32 `gorm:"not_null"`
	AnimalId  int
}
