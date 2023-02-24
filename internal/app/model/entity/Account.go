package entity

type Account struct {
	Id        int    `gorm:"primary_key"`
	FirstName string `gorm:"not_null"`
	LastName  string `gorm:"not_null"`
	Email     string `gorm:"not_null"`
	Password  string `gorm:"not_null"`
}
