package handler

import (
	"effective_mobile_rest/internal/api/v1/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "effective_mobile_rest/docs"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		cars := api.Group("/cars")
		{
			cars.POST("/", h.createCars)
			cars.DELETE("/:id", h.deleteCarById)
			cars.PATCH("/:id", h.updateCar)
			cars.GET("/", h.getAllCars)
		}
	}

	return router
}
