package router

import (
	"assignment3-012/main-service/internal/controller"
	"assignment3-012/main-service/internal/database"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	database.StartDB()
	router := gin.Default()

	router.PUT("/weather", controller.UpdateWeather)

	return router
}