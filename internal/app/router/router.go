package router

import (
	"github.com/gin-gonic/gin"
	"it-planet-task/helpers"
	"it-planet-task/internal/app/handler"
	"it-planet-task/internal/app/repository"
	"it-planet-task/internal/app/service"
	"it-planet-task/internal/app/service/geometry"
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

	animalLocationRepo := repository.NewAnimalLocationRepository(helpers.GetConnectionOrCreateAndGet(), animalRepo)
	animalLocationService := service.NewAnimalLocationService(animalLocationRepo)

	geometryService := geometry.NewGeometryService()

	areaRepo := repository.NewAreaRepository(helpers.GetConnectionOrCreateAndGet())
	areaService := service.NewAreaService(areaRepo, animalLocationService, geometryService)

	animalHandler := handler.NewAnimalHandler(animalService, animalTypeService, accountService, locationService, animalLocationService)
	animalGroup := api.Group("animals")
	{
		animalGroup.GET("/:id", middleware.BasicAuth, animalHandler.Get)
		animalGroup.GET("/search", middleware.BasicAuth, animalHandler.Search)
		animalGroup.POST("", middleware.BasicAuth, animalHandler.Create)
		animalGroup.PUT("/:id", middleware.BasicAuth, animalHandler.Update)
		animalGroup.DELETE("/:id", middleware.BasicAuth, middleware.AdminRequired, animalHandler.Delete)

		animalGroup.POST("/:id/types/:typeId", middleware.BasicAuth, middleware.AdminOrChipperRequired, animalHandler.AddAnimalType)
		animalGroup.PUT("/:id/types", middleware.BasicAuth, middleware.AdminOrChipperRequired, animalHandler.EditAnimalType)
		animalGroup.DELETE("/:id/types/:typeId", middleware.BasicAuth, middleware.AdminOrChipperRequired, animalHandler.DeleteAnimalType)
	}

	animalLocationHandler := handler.NewAnimalLocationHandler(animalLocationService, animalService, locationService)
	{
		animalGroup.GET("/:id/locations", middleware.BasicAuth, animalLocationHandler.GetAnimalLocations)
		animalGroup.POST("/:id/locations/:pointId", middleware.BasicAuth, middleware.AdminOrChipperRequired, animalLocationHandler.AddAnimalLocationPoint)
		animalGroup.PUT("/:id/locations", middleware.BasicAuth, middleware.AdminOrChipperRequired, animalLocationHandler.EditAnimalLocationPoint)
		animalGroup.DELETE("/:id/locations/:visitedPointId", middleware.BasicAuth, middleware.AdminRequired, animalLocationHandler.DeleteAnimalLocationPoint)
	}

	animalTypeHandler := handler.NewAnimalTypeHandler(animalTypeService, animalService)
	animalTypeGroup := animalGroup.Group("types")
	{
		animalTypeGroup.GET("/:id", middleware.BasicAuth, animalTypeHandler.Get)
		animalTypeGroup.POST("", middleware.BasicAuth, middleware.AdminOrChipperRequired, animalTypeHandler.Create)
		animalTypeGroup.PUT("/:id", middleware.BasicAuth, middleware.AdminOrChipperRequired, animalTypeHandler.Update)
		animalTypeGroup.DELETE("/:id", middleware.BasicAuth, middleware.AdminRequired, animalTypeHandler.Delete)
	}

	accountHandler := handler.NewAccountHandler(accountService, animalService)
	accountGroup := api.Group("accounts")
	{
		accountGroup.GET("/:id", middleware.BasicAuth, accountHandler.Get)
		accountGroup.GET("/search", middleware.BasicAuth, middleware.AdminRequired, accountHandler.Search)
		accountGroup.PUT("/:id", middleware.BasicAuth, accountHandler.Update)
		accountGroup.DELETE("/:id", middleware.BasicAuth, accountHandler.Delete)
		accountGroup.POST("", middleware.BasicAuth, middleware.AdminRequired, accountHandler.Create)
	}

	locationHandler := handler.NewLocationHandler(locationService, animalService)
	locationGroup := api.Group("locations")
	{
		locationGroup.GET("/:id", middleware.BasicAuth, locationHandler.Get)
		locationGroup.GET("", middleware.BasicAuth, locationHandler.GetByCoordinates)
		locationGroup.GET("/geohash", middleware.BasicAuth, locationHandler.GeoHashV1)
		locationGroup.GET("/geohashv2", middleware.BasicAuth, locationHandler.GeoHashV2)
		locationGroup.GET("/geohashv3", middleware.BasicAuth, locationHandler.GeoHashV3)
		locationGroup.POST("", middleware.BasicAuth, middleware.AdminOrChipperRequired, locationHandler.Create)
		locationGroup.PUT("/:id", middleware.BasicAuth, middleware.AdminOrChipperRequired, locationHandler.Update)
		locationGroup.DELETE("/:id", middleware.BasicAuth, middleware.AdminRequired, locationHandler.Delete)
	}

	authRepo := repository.NewAuthRepository(helpers.GetConnectionOrCreateAndGet())
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService, accountService)
	{
		r.POST("api/registration", authHandler.Register)
	}

	areaHandler := handler.NewAreaHandler(areaService, areaRepo)
	areaGroup := api.Group("areas")
	{
		areaGroup.GET("/:id", middleware.BasicAuth, areaHandler.Get)
		areaGroup.POST("", middleware.BasicAuth, middleware.AdminRequired, areaHandler.Create)
		areaGroup.PUT("/:id", middleware.BasicAuth, middleware.AdminRequired, areaHandler.Update)
		areaGroup.DELETE("/:id", middleware.BasicAuth, middleware.AdminRequired, areaHandler.Delete)
		areaGroup.GET("/:id/analytics", middleware.BasicAuth, areaHandler.Analytics)
	}

	return r
}
