package entity

const (
	UserRole    = "USER"
	ChipperRole = "CHIPPER"
	AdminRole   = "ADMIN"
)

type Account struct {
	Id        int    `gorm:"primary_key"`
	FirstName string `gorm:"not_null"`
	LastName  string `gorm:"not_null"`
	Email     string `gorm:"not_null"`
	Password  string `gorm:"not_null"`
	Role      string `gorm:"not_null"`
}
