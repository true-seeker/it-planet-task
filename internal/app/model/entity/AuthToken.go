package entity

type AuthToken struct {
	AccountId int    `gorm:"primaryKey" json:"accountId"`
	Token     string `json:"token"`
}
