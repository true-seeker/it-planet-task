package helpers

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
	"log"
	"time"
)

// GormMigrate Запуск миграций БД
func GormMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&entity.AnimalType{}, &entity.Account{}, &entity.Animal{}, &entity.Location{}, &entity.AnimalLocation{}, &entity.AuthToken{})
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

func floatToPtr(i float64) *float64 {
	return &i
}

func InitData(db *gorm.DB) {
	a := []entity.Location{
		{
			Latitude:  floatToPtr(50),
			Longitude: floatToPtr(100),
		},
		{
			Latitude:  floatToPtr(55),
			Longitude: floatToPtr(105),
		},
		{
			Latitude:  floatToPtr(60),
			Longitude: floatToPtr(110),
		},
		{
			Latitude:  floatToPtr(65),
			Longitude: floatToPtr(115),
		},
		{
			Latitude:  floatToPtr(70),
			Longitude: floatToPtr(120),
		},
		{
			Latitude:  floatToPtr(75),
			Longitude: floatToPtr(125),
		},
	}
	db.Save(&a)

	b := []entity.AnimalType{
		{
			Type: "Кот",
		},
		{
			Type: "Не кот",
		},
		{
			Type: "Собака",
		},
		{
			Type: "Капибара",
		},
	}
	db.Save(&b)

	c := []entity.Animal{
		{
			AnimalTypes:        []entity.AnimalType{b[0]},
			Weight:             1,
			Height:             2,
			Length:             3,
			Gender:             "FEMALE",
			LifeStatus:         "ALIVE",
			ChippingDateTime:   time.Time{},
			ChipperId:          1,
			ChippingLocationId: 1,
		},
		{
			AnimalTypes:        []entity.AnimalType{b[1]},
			Weight:             1,
			Height:             2,
			Length:             3,
			Gender:             "FEMALE",
			LifeStatus:         "ALIVE",
			ChippingDateTime:   time.Time{},
			ChipperId:          1,
			ChippingLocationId: 1,
		},
		{
			AnimalTypes:        []entity.AnimalType{b[2]},
			Weight:             1,
			Height:             2,
			Length:             3,
			Gender:             "FEMALE",
			LifeStatus:         "ALIVE",
			ChippingDateTime:   time.Time{},
			ChipperId:          1,
			ChippingLocationId: 1,
		},
		{
			AnimalTypes:        []entity.AnimalType{b[3]},
			Weight:             1,
			Height:             2,
			Length:             3,
			Gender:             "FEMALE",
			LifeStatus:         "ALIVE",
			ChippingDateTime:   time.Time{},
			ChipperId:          1,
			ChippingLocationId: 1,
		},
		{
			AnimalTypes:        []entity.AnimalType{b[0]},
			Weight:             1,
			Height:             2,
			Length:             3,
			Gender:             "FEMALE",
			LifeStatus:         "ALIVE",
			ChippingDateTime:   time.Time{},
			ChipperId:          1,
			ChippingLocationId: 1,
		},
		{
			AnimalTypes:        []entity.AnimalType{b[1]},
			Weight:             1,
			Height:             2,
			Length:             3,
			Gender:             "FEMALE",
			LifeStatus:         "ALIVE",
			ChippingDateTime:   time.Time{},
			ChipperId:          1,
			ChippingLocationId: 1,
		},
		{
			AnimalTypes:        []entity.AnimalType{b[2]},
			Weight:             1,
			Height:             2,
			Length:             3,
			Gender:             "FEMALE",
			LifeStatus:         "ALIVE",
			ChippingDateTime:   time.Time{},
			ChipperId:          1,
			ChippingLocationId: 1,
		},
		{
			AnimalTypes:        []entity.AnimalType{b[3]},
			Weight:             1,
			Height:             2,
			Length:             3,
			Gender:             "FEMALE",
			LifeStatus:         "ALIVE",
			ChippingDateTime:   time.Time{},
			ChipperId:          1,
			ChippingLocationId: 1,
		},
		{
			AnimalTypes:        []entity.AnimalType{b[0]},
			Weight:             1,
			Height:             2,
			Length:             3,
			Gender:             "FEMALE",
			LifeStatus:         "ALIVE",
			ChippingDateTime:   time.Time{},
			ChipperId:          1,
			ChippingLocationId: 1,
		},
		{
			AnimalTypes:        []entity.AnimalType{b[1]},
			Weight:             1,
			Height:             2,
			Length:             3,
			Gender:             "FEMALE",
			LifeStatus:         "ALIVE",
			ChippingDateTime:   time.Time{},
			ChipperId:          1,
			ChippingLocationId: 1,
		},
		{
			AnimalTypes:        []entity.AnimalType{b[2]},
			Weight:             1,
			Height:             2,
			Length:             3,
			Gender:             "FEMALE",
			LifeStatus:         "ALIVE",
			ChippingDateTime:   time.Time{},
			ChipperId:          1,
			ChippingLocationId: 1,
		},
		{
			AnimalTypes:        []entity.AnimalType{b[3]},
			Weight:             1,
			Height:             2,
			Length:             3,
			Gender:             "FEMALE",
			LifeStatus:         "ALIVE",
			ChippingDateTime:   time.Time{},
			ChipperId:          1,
			ChippingLocationId: 1,
		},
	}
	db.Save(&c)
}
