package router

import (
	"github.com/gin-gonic/gin"
	"it-planet-task/helpers"
	"it-planet-task/internal/app/handler"
	"it-planet-task/internal/app/repository"
	"it-planet-task/internal/app/service"
)

func New(r *gin.Engine) *gin.Engine {
	NewAnimalTypeHandler(r)
	NewAccountHandler(r)
	NewLocationHandler(r)

	return r
}

func NewAnimalTypeHandler(r *gin.Engine) {
	animalTypeRepo := repository.NewAnimalTypeRepository(helpers.GetConnectionOrCreateAndGet())
	animalTypeService := service.NewAnimalTypeService(animalTypeRepo)
	h := handler.NewAnimalTypeHandler(animalTypeService)
	animalTypeGroup := r.Group("animal_type")

	animalTypeGroup.GET("/:id", h.Get)

}

func NewAccountHandler(r *gin.Engine) {
	accountRepo := repository.NewAccountRepository(helpers.GetConnectionOrCreateAndGet())
	accountService := service.NewAccountService(accountRepo)
	h := handler.NewAccountHandler(accountService)
	accountGroup := r.Group("accounts")

	accountGroup.GET("/:id", h.Get)
}

func NewLocationHandler(r *gin.Engine) {
	locationRepo := repository.NewLocationRepository(helpers.GetConnectionOrCreateAndGet())
	locationService := service.NewLocationService(locationRepo)
	h := handler.NewLocationHandler(locationService)
	locationGroup := r.Group("locations")

	locationGroup.GET("/:id", h.Get)
}
