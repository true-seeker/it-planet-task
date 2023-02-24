package entity

type AnimalType struct {
	Id    int    `gorm:"primary_key"`
	Title string `gorm:"not_null"`
}
