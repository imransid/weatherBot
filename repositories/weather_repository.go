package repositories

import (
	"database/sql"
	"fmt"

	"example.com/weatherApp/models"
)

type WeatherRepository interface {
	GetWeatherByCity(city string) (*models.Weather, error)
	SaveWeather(weather *models.Weather) error
}

type weatherRepository struct {
	db *sql.DB
}

func NewWeatherRepository(db *sql.DB) WeatherRepository {
	return &weatherRepository{
		db: db,
	}
}

func (r *weatherRepository) GetWeatherByCity(city string) (*models.Weather, error) {
	query := "SELECT id, city, temperature, conditions FROM weather WHERE city = $1"
	row := r.db.QueryRow(query, city)

	weather := &models.Weather{}
	err := row.Scan(&weather.ID, &weather.City, &weather.Temperature, &weather.Conditions)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("weather not found for city: %s", city)
		}
		return nil, fmt.Errorf("failed to get weather: %w", err)
	}

	return weather, nil
}

func (r *weatherRepository) SaveWeather(weather *models.Weather) error {
	query := "INSERT INTO weather (city, temperature, conditions) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, weather.City, weather.Temperature, weather.Conditions).Scan(&weather.ID)
	if err != nil {
		return fmt.Errorf("failed to save weather: %w", err)
	}

	return nil
}
