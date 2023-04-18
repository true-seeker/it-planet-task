package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"it-planet-task/helpers"
	"it-planet-task/internal/pkg/app"
	"it-planet-task/pkg/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Init Инициализация сервиса
func Init() {
	config.Init("development")
	helpers.GormMigrate(helpers.GetConnectionOrCreateAndGet())
	helpers.InitAccounts(helpers.GetConnectionOrCreateAndGet())
}

func main() {
	Init()
	r := gin.Default()

	a := app.New(r)
	srv := a.GetServer()

	go func() {
		err := a.Run(srv)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Closing DB connection")
	db := helpers.GetConnectionOrCreateAndGet()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("sqlDB:", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatal("sqlDB close:", err)
	}

	log.Println("Stopping http server")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}

	select {
	case <-ctx.Done():
	}
}
