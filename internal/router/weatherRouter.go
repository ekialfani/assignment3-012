package router

import (
	"assignment3-012/internal/controller"
	"assignment3-012/internal/database"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	database.StartDB()
	router := gin.Default()

	router.PUT("/weather", controller.UpdateWeather)

	return router
}