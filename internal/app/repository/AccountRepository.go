package repository

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
)

type Account interface {
	Get(id int) (*entity.Account, error)
	Search() (*[]entity.Account, error)
}

type AccountRepository struct {
	Db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) Account {
	return &AccountRepository{Db: db}
}

func (a *AccountRepository) Get(id int) (*entity.Account, error) {
	var account entity.Account
	err := a.Db.First(&account, id).Error
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *AccountRepository) Search() (*[]entity.Account, error) {
	var accounts []entity.Account
	err := a.Db.Find(&accounts).Error
	if err != nil {
		return nil, err
	}

	return &accounts, nil
}