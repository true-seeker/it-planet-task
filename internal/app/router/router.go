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
	return r
}

func NewAnimalTypeHandler(r *gin.Engine) {
	animalTypeRepo := repository.NewAnimalTypeRepository(helpers.GetConnectionOrCreateAndGet())
	animalTypeService := service.NewAnimalTypeService(animalTypeRepo)
	h := handler.NewAnimalTypeHandler(animalTypeService)
	animalTypeGroup := r.Group("animal_type")

	animalTypeGroup.GET("/:id", h.Get)

}
