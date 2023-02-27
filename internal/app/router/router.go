package router

import (
	"github.com/gin-gonic/gin"
	"it-planet-task/helpers"
	"it-planet-task/internal/app/handler"
	"it-planet-task/internal/app/repository"
	"it-planet-task/internal/app/service"
)

func New(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")

	animalRepo := repository.NewAnimalRepository(helpers.GetConnectionOrCreateAndGet())
	animalService := service.NewAnimalService(animalRepo)
	animalHandler := handler.NewAnimalHandler(animalService)
	animalGroup := api.Group("animals")
	{
		animalGroup.GET("/:id", animalHandler.Get)
		animalGroup.GET("/:id/locations", animalHandler.GetAnimalLocations)
		animalGroup.GET("/search", animalHandler.Search)
	}

	animalTypeRepo := repository.NewAnimalTypeRepository(helpers.GetConnectionOrCreateAndGet())
	animalTypeService := service.NewAnimalTypeService(animalTypeRepo)
	animalTypeHandler := handler.NewAnimalTypeHandler(animalTypeService)
	animalTypeGroup := animalGroup.Group("types")
	{
		animalTypeGroup.GET("/:id", animalTypeHandler.Get)
	}

	accountRepo := repository.NewAccountRepository(helpers.GetConnectionOrCreateAndGet())
	accountService := service.NewAccountService(accountRepo)
	accountHandler := handler.NewAccountHandler(accountService)
	accountGroup := api.Group("accounts")
	{
		accountGroup.GET("/:id", accountHandler.Get)
		accountGroup.GET("/search", accountHandler.Search)
	}

	locationRepo := repository.NewLocationRepository(helpers.GetConnectionOrCreateAndGet())
	locationService := service.NewLocationService(locationRepo)
	locationHandler := handler.NewLocationHandler(locationService)
	locationGroup := api.Group("locations")
	{
		locationGroup.GET("/:id", locationHandler.Get)
	}

	authRepo := repository.NewAuthRepository(helpers.GetConnectionOrCreateAndGet())
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService, accountService)
	{
		r.POST("/registration", authHandler.Register)
	}

	return r
}
