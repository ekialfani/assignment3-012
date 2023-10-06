package controller

import (
	"assignment3-012/internal/models"
	"assignment3-012/internal/service"
	"fmt"
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

	fmt.Println(weather)
	fmt.Println("status water:", waterStatus)
	fmt.Println("status wind:", windStatus)
}