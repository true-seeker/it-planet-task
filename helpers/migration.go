package helpers

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/model/entity"
	"log"
)

func GormMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&entity.AnimalType{})
	if err != nil {
		log.Fatal(err)
	}
}
