package service

import (
	"assignment3-012/internal/models"
	"assignment3-012/internal/repository"
)

func UpdateWeather(weather *models.Weather) error {
	totalWeatherData := repository.GetTotalWeatherData()

	if totalWeatherData == 0 {
		return repository.CreateWeather(weather)
	}

	return repository.UpdateWeather(weather)
}

func GetWaterStatus(water int) string {
	if water <= 5 {
		return "aman"
	}

	if water >= 6 && water <= 8 {
		return "siaga"
	}

	return "bahaya"
}

func GetWindStatus(wind int) string {
	if wind <= 6 {
		return "aman"
	}

	if wind >= 7 && wind <= 15 {
		return "siaga"
	}

	return "bahaya"
}