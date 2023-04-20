package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"it-planet-task/helpers"
	"it-planet-task/internal/pkg/app"
	"it-planet-task/pkg/config"
	"log"
)

// Init Инициализация сервиса
func Init() {
	config.Init("development")
	helpers.GormMigrate(helpers.GetConnectionOrCreateAndGet())
	helpers.InitAccounts(helpers.GetConnectionOrCreateAndGet())
	helpers.InitData(helpers.GetConnectionOrCreateAndGet())
}

func main() {
	Init()
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	a := app.New(r)
	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
