package router

import (
	"github.com/gin-gonic/gin"
	"it-planet-task/helpers"
	"it-planet-task/internal/app/handler"
	"it-planet-task/internal/app/repository"
	"it-planet-task/internal/app/service"
)

func New(r *gin.Engine) *gin.Engine {
	animalGroup := NewAnimalRouter(r)
	NewAnimalTypeRouter(animalGroup)
	NewAccountRouter(r)
	NewLocationRouter(r)
	NewAuthRouter(r)

	return r
}

func NewAnimalTypeRouter(parentGroup *gin.RouterGroup) {
	animalTypeRepo := repository.NewAnimalTypeRepository(helpers.GetConnectionOrCreateAndGet())
	animalTypeService := service.NewAnimalTypeService(animalTypeRepo)
	h := handler.NewAnimalTypeHandler(animalTypeService)
	animalTypeGroup := parentGroup.Group("types")

	animalTypeGroup.GET("/:id", h.Get)

}

func NewAccountRouter(r *gin.Engine) {
	accountRepo := repository.NewAccountRepository(helpers.GetConnectionOrCreateAndGet())
	accountService := service.NewAccountService(accountRepo)
	h := handler.NewAccountHandler(accountService)
	accountGroup := r.Group("accounts")

	accountGroup.GET("/:id", h.Get)
	accountGroup.GET("/search", h.Search)
}

func NewLocationRouter(r *gin.Engine) {
	locationRepo := repository.NewLocationRepository(helpers.GetConnectionOrCreateAndGet())
	locationService := service.NewLocationService(locationRepo)
	h := handler.NewLocationHandler(locationService)
	locationGroup := r.Group("locations")

	locationGroup.GET("/:id", h.Get)
}

func NewAnimalRouter(r *gin.Engine) *gin.RouterGroup {
	animalRepo := repository.NewAnimalRepository(helpers.GetConnectionOrCreateAndGet())
	animalService := service.NewAnimalService(animalRepo)
	h := handler.NewAnimalHandler(animalService)
	animalGroup := r.Group("animals")

	animalGroup.GET("/:id", h.Get)
	animalGroup.GET("/:id/locations", h.GetAnimalLocations)
	animalGroup.GET("/search", h.Search)
	return animalGroup
}

func NewAuthRouter(r *gin.Engine) {
	authRepo := repository.NewAuthRepository(helpers.GetConnectionOrCreateAndGet())
	authService := service.NewAuthService(authRepo)
	accountRepo := repository.NewAccountRepository(helpers.GetConnectionOrCreateAndGet())
	accountService := service.NewAccountService(accountRepo)
	h := handler.NewAuthHandler(authService, accountService)

	r.POST("/registration", h.Register)
}
