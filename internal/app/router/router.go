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
		animalGroup.GET("/:id", middleware.TokenAuth, animalHandler.Get)
		animalGroup.GET("/search", middleware.TokenAuth, animalHandler.Search)
		animalGroup.POST("", middleware.TokenAuth, animalHandler.Create)
		animalGroup.PUT("/:id", middleware.TokenAuth, animalHandler.Update)
		animalGroup.DELETE("/:id", middleware.TokenAuth, middleware.AdminRequired, animalHandler.Delete)

		animalGroup.POST("/:id/types/:typeId", middleware.TokenAuth, animalHandler.AddAnimalType)
		animalGroup.PUT("/:id/types", middleware.TokenAuth, animalHandler.EditAnimalType)
		animalGroup.DELETE("/:id/types/:typeId", middleware.TokenAuth, animalHandler.DeleteAnimalType)
	}

	animalLocationHandler := handler.NewAnimalLocationHandler(animalLocationService, animalService, locationService)
	{
		animalGroup.GET("/:id/locations", middleware.TokenAuth, animalLocationHandler.GetAnimalLocations)
		animalGroup.POST("/:id/locations/:pointId", middleware.TokenAuth, middleware.AdminOrChipperRequired, animalLocationHandler.AddAnimalLocationPoint)
		animalGroup.PUT("/:id/locations", middleware.TokenAuth, middleware.AdminOrChipperRequired, animalLocationHandler.EditAnimalLocationPoint)
		animalGroup.DELETE("/:id/locations/:visitedPointId", middleware.TokenAuth, middleware.AdminRequired, animalLocationHandler.DeleteAnimalLocationPoint)
	}

	animalTypeHandler := handler.NewAnimalTypeHandler(animalTypeService, animalService)
	animalTypeGroup := animalGroup.Group("types")
	{
		animalTypeGroup.GET("/:id", middleware.TokenAuth, animalTypeHandler.Get)
		animalTypeGroup.POST("", middleware.TokenAuth, animalTypeHandler.Create)
		animalTypeGroup.PUT("/:id", middleware.TokenAuth, animalTypeHandler.Update)
		animalTypeGroup.DELETE("/:id", middleware.TokenAuth, animalTypeHandler.Delete)
		animalTypeGroup.GET("/search", middleware.TokenAuth, animalTypeHandler.Search)
	}

	accountHandler := handler.NewAccountHandler(accountService, animalService)
	accountGroup := api.Group("accounts")
	{
		accountGroup.GET("/:id", middleware.TokenAuth, accountHandler.Get)
		accountGroup.GET("/search", middleware.TokenAuth, middleware.AdminRequired, accountHandler.Search)
		accountGroup.PUT("/:id", middleware.TokenAuth, accountHandler.Update)
		accountGroup.DELETE("/:id", middleware.TokenAuth, accountHandler.Delete)
		accountGroup.POST("", middleware.TokenAuth, middleware.AdminRequired, accountHandler.Create)
	}

	locationHandler := handler.NewLocationHandler(locationService, animalService)
	locationGroup := api.Group("locations")
	{
		locationGroup.GET("/:id", middleware.TokenAuth, locationHandler.Get)
		locationGroup.POST("", middleware.TokenAuth, middleware.AdminOrChipperRequired, locationHandler.Create)
		locationGroup.PUT("/:id", middleware.TokenAuth, middleware.AdminOrChipperRequired, locationHandler.Update)
		locationGroup.DELETE("/:id", middleware.TokenAuth, middleware.AdminRequired, locationHandler.Delete)
	}

	authRepo := repository.NewAuthRepository(helpers.GetConnectionOrCreateAndGet())
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService, accountService)
	{
		r.POST("api/registration", authHandler.Register)
		r.POST("api/login", authHandler.Login)
	}

	mockHandler := handler.NewMockHandler(helpers.GetConnectionOrCreateAndGet())
	r.GET("api/generate_data", mockHandler.GenerateRandomData)

	return r
}
