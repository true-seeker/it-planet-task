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

	animalTypeRepo := repository.NewAnimalTypeRepository(helpers.GetConnectionOrCreateAndGet())
	animalTypeService := service.NewAnimalTypeService(animalTypeRepo)

	accountRepo := repository.NewAccountRepository(helpers.GetConnectionOrCreateAndGet())
	accountService := service.NewAccountService(accountRepo)

	locationRepo := repository.NewLocationRepository(helpers.GetConnectionOrCreateAndGet())
	locationService := service.NewLocationService(locationRepo)

	animalRepo := repository.NewAnimalRepository(helpers.GetConnectionOrCreateAndGet())
	animalService := service.NewAnimalService(animalRepo)

	animalHandler := handler.NewAnimalHandler(animalService, animalTypeService, accountService, locationService)
	animalGroup := api.Group("animals")
	{
		animalGroup.GET("/:id", middleware.OptionalBasicAuth, animalHandler.Get)
		animalGroup.GET("/:id/locations", middleware.OptionalBasicAuth, animalHandler.GetAnimalLocations)
		animalGroup.GET("/search", middleware.OptionalBasicAuth, animalHandler.Search)
		animalGroup.POST("/", middleware.BasicAuth, animalHandler.Create)
		animalGroup.PUT("/:id", middleware.BasicAuth, animalHandler.Update)
		animalGroup.DELETE("/:id", middleware.BasicAuth, animalHandler.Delete)
		animalGroup.POST("/:animalId/types/:typeId", middleware.BasicAuth, animalHandler.AddAnimalType)
		animalGroup.PUT("/:id/types", middleware.BasicAuth, animalHandler.EditAnimalType)
		animalGroup.DELETE("/:animalId/types/:typeId", middleware.BasicAuth, animalHandler.DeleteAnimalType)
	}

	animalTypeHandler := handler.NewAnimalTypeHandler(animalTypeService, animalService)
	animalTypeGroup := animalGroup.Group("types")
	{
		animalTypeGroup.GET("/:id", middleware.OptionalBasicAuth, animalTypeHandler.Get)
		animalTypeGroup.POST("/", middleware.BasicAuth, animalTypeHandler.Create)
		animalTypeGroup.PUT("/:id", middleware.BasicAuth, animalTypeHandler.Update)
		animalTypeGroup.DELETE("/:id", middleware.BasicAuth, animalTypeHandler.Delete)
	}

	accountHandler := handler.NewAccountHandler(accountService, animalService)
	accountGroup := api.Group("accounts")
	{
		accountGroup.GET("/:id", middleware.OptionalBasicAuth, accountHandler.Get)
		accountGroup.GET("/search", middleware.OptionalBasicAuth, accountHandler.Search)
		accountGroup.PUT("/:id", middleware.BasicAuth, accountHandler.Update)
		accountGroup.DELETE("/:id", middleware.BasicAuth, accountHandler.Delete)
	}

	locationHandler := handler.NewLocationHandler(locationService, animalService)
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
