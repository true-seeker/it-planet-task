package repository

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/filter"
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/pkg/paginator"
)

type Account interface {
	Get(id int) (*entity.Account, error)
	Update(account *entity.Account) (*entity.Account, error)
	Search(params *filter.AccountFilterParams) (*[]entity.Account, error)
	GetByEmail(account *entity.Account) *entity.Account
	CheckCredentials(account *entity.Account) bool
	Delete(id int) error
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

func (a *AccountRepository) Search(params *filter.AccountFilterParams) (*[]entity.Account, error) {
	var accounts []entity.Account
	err := a.Db.
		Scopes(paginator.Paginate(params),
			filter.AccountFilter(params)).
		Order("id").
		Find(&accounts).Error
	if err != nil {
		return nil, err
	}

	return &accounts, nil
}

func (a *AccountRepository) GetByEmail(account *entity.Account) *entity.Account {
	ac := &entity.Account{}
	a.Db.Where("email = ?", account.Email).First(ac)
	return ac
}

func (a *AccountRepository) CheckCredentials(account *entity.Account) bool {
	var acc entity.Account
	a.Db.Where("email = ? AND password = ?", account.Email, account.Password).First(&acc)
	return acc.Id != 0
}

func (a *AccountRepository) Update(account *entity.Account) (*entity.Account, error) {
	err := a.Db.Save(&account).Error
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (a *AccountRepository) Delete(id int) error {
	err := a.Db.Delete(&entity.Account{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
