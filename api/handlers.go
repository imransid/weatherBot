package api

import (
	"net/http"

	"example.com/weatherApp/repositories"
	"github.com/gin-gonic/gin"
)

type WeatherAPI struct {
	WeatherRepo repositories.WeatherRepository
}

func (api *WeatherAPI) GetWeather(c *gin.Context) {
	city := c.Param("city")

	weatherData, err := api.WeatherRepo.GetWeatherByCity(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weather data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"city":        weatherData.City,
		"temperature": weatherData.Temperature,
		"conditions":  weatherData.Conditions,
	})
}

func InitRoutes(router *gin.Engine, weatherRepo repositories.WeatherRepository) {
	weatherAPI := &WeatherAPI{
		WeatherRepo: weatherRepo,
	}

	router.GET("/weather/:city", weatherAPI.GetWeather)
}
