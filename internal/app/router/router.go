package router

import (
	"github.com/gin-gonic/gin"
	"it-planet-task/helpers"
	"it-planet-task/internal/app/handler"
	"it-planet-task/internal/app/repository"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/pkg/middleware"
)

func New(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")

	animalRepo := repository.NewAnimalRepository(helpers.GetConnectionOrCreateAndGet())
	animalService := service.NewAnimalService(animalRepo)
	animalHandler := handler.NewAnimalHandler(animalService)
	animalGroup := api.Group("animals")
	{
		animalGroup.GET("/:id", middleware.OptionalBasicAuth, animalHandler.Get)
		animalGroup.GET("/:id/locations", middleware.OptionalBasicAuth, animalHandler.GetAnimalLocations)
		animalGroup.GET("/search", middleware.OptionalBasicAuth, animalHandler.Search)
	}

	animalTypeRepo := repository.NewAnimalTypeRepository(helpers.GetConnectionOrCreateAndGet())
	animalTypeService := service.NewAnimalTypeService(animalTypeRepo)
	animalTypeHandler := handler.NewAnimalTypeHandler(animalTypeService, animalService)
	animalTypeGroup := animalGroup.Group("types")
	{
		animalTypeGroup.GET("/:id", middleware.OptionalBasicAuth, animalTypeHandler.Get)
		animalTypeGroup.POST("/", middleware.BasicAuth, animalTypeHandler.Create)
		animalTypeGroup.PUT("/:id", middleware.BasicAuth, animalTypeHandler.Update)
		animalTypeGroup.DELETE("/:id", middleware.BasicAuth, animalTypeHandler.Delete)
	}

	accountRepo := repository.NewAccountRepository(helpers.GetConnectionOrCreateAndGet())
	accountService := service.NewAccountService(accountRepo)
	accountHandler := handler.NewAccountHandler(accountService, animalService)
	accountGroup := api.Group("accounts")
	{
		accountGroup.GET("/:id", middleware.OptionalBasicAuth, accountHandler.Get)
		accountGroup.GET("/search", middleware.OptionalBasicAuth, accountHandler.Search)
		accountGroup.PUT("/:id", middleware.BasicAuth, accountHandler.Update)
		accountGroup.DELETE("/:id", middleware.BasicAuth, accountHandler.Delete)
	}

	locationRepo := repository.NewLocationRepository(helpers.GetConnectionOrCreateAndGet())
	locationService := service.NewLocationService(locationRepo)
	locationHandler := handler.NewLocationHandler(locationService)
	locationGroup := api.Group("locations")
	{
		locationGroup.GET("/:id", middleware.OptionalBasicAuth, locationHandler.Get)
		locationGroup.POST("/", middleware.BasicAuth, locationHandler.Create)
		locationGroup.PUT("/:id", middleware.BasicAuth, locationHandler.Update)
		locationGroup.DELETE("/:id", middleware.BasicAuth, locationHandler.Delete)
	}

	authRepo := repository.NewAuthRepository(helpers.GetConnectionOrCreateAndGet())
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService, accountService)
	{
		r.POST("api/registration", authHandler.Register)
	}

	return r
}
