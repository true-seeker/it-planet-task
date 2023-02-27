package repository

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
)

type Auth interface {
	Register(newAccount *entity.Account) (*entity.Account, error)
}

type AuthRepository struct {
	Db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) Auth {
	return &AuthRepository{Db: db}
}

func (a AuthRepository) Register(newAccount *entity.Account) (*entity.Account, error) {
	err := a.Db.Save(&newAccount).Error
	if err != nil {
		return nil, err
	}

	return newAccount, nil
}
