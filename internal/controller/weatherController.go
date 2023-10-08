package controller

import (
	"assignment3-012/internal/models"
	"assignment3-012/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateWeather(context *gin.Context) {
	var weather models.Weather

	err := context.ShouldBindJSON(&weather)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	err = service.UpdateWeather(&weather)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	waterStatus := service.GetWaterStatus(weather.Water)
	windStatus := service.GetWindStatus(weather.Wind)

	context.JSON(http.StatusOK, gin.H{
		"weather": weather,
		"water_status": waterStatus,
		"wind_status": windStatus,
	})
}