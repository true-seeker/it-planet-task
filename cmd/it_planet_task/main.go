package main

import (
	"github.com/gin-gonic/gin"
	"it-planet-task/helpers"
	"it-planet-task/internal/pkg/app"
	"it-planet-task/pkg/config"
	"log"
)

func Init() {
	config.Init("development")
	helpers.GormMigrate(helpers.GetConnectionOrCreateAndGet())
}

func main() {
	Init()
	r := gin.Default()
	a := app.New(r)

	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
