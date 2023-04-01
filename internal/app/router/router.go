package router

import (
	"github.com/gin-gonic/gin"
	"it-planet-task/helpers"
	"it-planet-task/internal/app/handler"
	"it-planet-task/internal/app/repository"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/pkg/middleware"
)

// InitRoutes Инициализация путей эндпоинтов, сервисов и репозиториев
func InitRoutes(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")

	animalTypeRepo := repository.NewAnimalTypeRepository(helpers.GetConnectionOrCreateAndGet())
	animalTypeService := service.NewAnimalTypeService(animalTypeRepo)

	accountRepo := repository.NewAccountRepository(helpers.GetConnectionOrCreateAndGet())
	accountService := service.NewAccountService(accountRepo)

	locationRepo := repository.NewLocationRepository(helpers.GetConnectionOrCreateAndGet())
	locationService := service.NewLocationService(locationRepo)

	animalRepo := repository.NewAnimalRepository(helpers.GetConnectionOrCreateAndGet())
	animalService := service.NewAnimalService(animalRepo)

	animalLocationRepo := repository.NewAnimalLocationRepository(helpers.GetConnectionOrCreateAndGet())
	animalLocationService := service.NewAnimalLocationService(animalLocationRepo)

	animalHandler := handler.NewAnimalHandler(animalService, animalTypeService, accountService, locationService, animalLocationService)
	animalGroup := api.Group("animals")
	{
		animalGroup.GET("/:id", middleware.BasicAuth, animalHandler.Get)
		animalGroup.GET("/search", middleware.BasicAuth, animalHandler.Search)
		animalGroup.POST("", middleware.BasicAuth, animalHandler.Create)
		animalGroup.PUT("/:id", middleware.BasicAuth, animalHandler.Update)
		animalGroup.DELETE("/:id", middleware.BasicAuth, animalHandler.Delete)

		animalGroup.POST("/:id/types/:typeId", middleware.BasicAuth, animalHandler.AddAnimalType)
		animalGroup.PUT("/:id/types", middleware.BasicAuth, animalHandler.EditAnimalType)
		animalGroup.DELETE("/:id/types/:typeId", middleware.BasicAuth, animalHandler.DeleteAnimalType)
	}

	animalLocationHandler := handler.NewAnimalLocationHandler(animalLocationService, animalService, locationService)
	{
		animalGroup.GET("/:id/locations", middleware.BasicAuth, animalLocationHandler.GetAnimalLocations)
		animalGroup.POST("/:id/locations/:pointId", middleware.BasicAuth, animalLocationHandler.AddAnimalLocationPoint)
		animalGroup.PUT("/:id/locations", middleware.BasicAuth, animalLocationHandler.EditAnimalLocationPoint)
		animalGroup.DELETE("/:id/locations/:visitedPointId", middleware.BasicAuth, animalLocationHandler.DeleteAnimalLocationPoint)
	}

	animalTypeHandler := handler.NewAnimalTypeHandler(animalTypeService, animalService)
	animalTypeGroup := animalGroup.Group("types")
	{
		animalTypeGroup.GET("/:id", middleware.BasicAuth, animalTypeHandler.Get)
		animalTypeGroup.POST("", middleware.BasicAuth, animalTypeHandler.Create)
		animalTypeGroup.PUT("/:id", middleware.BasicAuth, animalTypeHandler.Update)
		animalTypeGroup.DELETE("/:id", middleware.BasicAuth, animalTypeHandler.Delete)
	}

	accountHandler := handler.NewAccountHandler(accountService, animalService)
	accountGroup := api.Group("accounts")
	{
		accountGroup.GET("/:id", middleware.BasicAuth, accountHandler.Get)
		accountGroup.GET("/search", middleware.BasicAuth, accountHandler.Search)
		accountGroup.PUT("/:id", middleware.BasicAuth, accountHandler.Update)
		accountGroup.DELETE("/:id", middleware.BasicAuth, accountHandler.Delete)
		accountGroup.POST("", middleware.BasicAuth, accountHandler.Create)
	}

	locationHandler := handler.NewLocationHandler(locationService, animalService)
	locationGroup := api.Group("locations")
	{
		locationGroup.GET("/:id", middleware.BasicAuth, locationHandler.Get)
		locationGroup.POST("", middleware.BasicAuth, locationHandler.Create)
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
