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
