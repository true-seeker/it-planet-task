package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"it-planet-task/internal/app/router"
	"it-planet-task/pkg/config"
)

type App struct {
	router *gin.Engine
}

func New(r *gin.Engine) *App {
	return &App{router: router.New(r)}
}

func (a *App) Run() error {
	err := a.router.Run(fmt.Sprintf("%s:%s", config.GetConfig().GetString("server.address"),
		config.GetConfig().GetString("server.port")))
	if err != nil {
		return err
	}
	return nil
}
