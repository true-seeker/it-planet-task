package helpers

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
	"log"
)

// GormMigrate Запуск миграций БД
func GormMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&entity.AnimalType{}, &entity.Account{}, &entity.Animal{}, &entity.Location{},
		&entity.AnimalLocation{}, &entity.Area{}, &entity.AreaPoint{})
	if err != nil {
		log.Fatal(err)
	}
}

// InitAccounts Инициализация начальных аккаунтов
func InitAccounts(db *gorm.DB) {
	adminAccount := &entity.Account{
		FirstName: "adminFirstName",
		LastName:  "adminLastName",
		Email:     "admin@simbirsoft.com",
		Password:  "qwerty123",
		Role:      "ADMIN",
	}

	chipperAccount := &entity.Account{
		FirstName: "chipperFirstName",
		LastName:  "chipperLastName",
		Email:     "chipper@simbirsoft.com",
		Password:  "qwerty123",
		Role:      "CHIPPER",
	}

	userAccount := &entity.Account{
		FirstName: "userFirstName",
		LastName:  "userLastName",
		Email:     "user@simbirsoft.com",
		Password:  "qwerty123",
		Role:      "USER",
	}

	db.Create(adminAccount)
	db.Create(chipperAccount)
	db.Create(userAccount)
}
