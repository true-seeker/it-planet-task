package helpers

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
	"log"
)

// GormMigrate Запуск миграций БД
func GormMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&entity.AnimalType{}, &entity.Account{}, &entity.Animal{}, &entity.Location{}, &entity.AnimalLocation{})
	if err != nil {
		log.Fatal(err)
	}
}

// InitAccounts Инициализация начальных аккаунтов
func InitAccounts(db *gorm.DB) {
	adminAccount := &entity.Account{
		Id:        1,
		FirstName: "adminFirstName",
		LastName:  "adminLastName",
		Email:     "admin@simbirsoft.com",
		Password:  "qwerty123",
		Role:      "ADMIN",
	}

	chipperAccount := &entity.Account{
		Id:        2,
		FirstName: "chipperFirstName",
		LastName:  "chipperLastName",
		Email:     "chipper@simbirsoft.com",
		Password:  "qwerty123",
		Role:      "CHIPPER",
	}

	userAccount := &entity.Account{
		Id:        3,
		FirstName: "userFirstName",
		LastName:  "userLastName",
		Email:     "user@simbirsoft.com",
		Password:  "qwerty123",
		Role:      "USER",
	}

	db.Save(adminAccount)
	db.Save(chipperAccount)
	db.Save(userAccount)
}
