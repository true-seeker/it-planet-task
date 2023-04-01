package entity

type Area struct {
	Id         int         `gorm:"primary_key"`
	Name       string      `gorm:"not_null"`
	AreaPoints []AreaPoint `gorm:"not_null"`
}
