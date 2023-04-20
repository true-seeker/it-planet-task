package repository

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
)

type Auth interface {
	Register(newAccount *entity.Account) (*entity.Account, error)
	SaveToken(authToken *entity.AuthToken) (*entity.AuthToken, error)
	CheckToken(authToken *entity.AuthToken) (*entity.AuthToken, error)
}

type AuthRepository struct {
	Db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) Auth {
	return &AuthRepository{Db: db}
}

func (a *AuthRepository) Register(newAccount *entity.Account) (*entity.Account, error) {
	err := a.Db.Save(&newAccount).Error
	if err != nil {
		return nil, err
	}

	return newAccount, nil
}

func (a *AuthRepository) SaveToken(authToken *entity.AuthToken) (*entity.AuthToken, error) {
	err := a.Db.Save(&authToken).Error
	if err != nil {
		return nil, err
	}

	return authToken, nil
}

func (a *AuthRepository) CheckToken(authToken *entity.AuthToken) (*entity.AuthToken, error) {
	err := a.Db.Where("token = ?", authToken.Token).First(&authToken).Error
	if err != nil {
		return nil, err
	}

	return authToken, nil
}
