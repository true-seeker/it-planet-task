package entity

type AnimalType struct {
	Id   int    `gorm:"primary_key"`
	Type string `gorm:"not_null"`
}
