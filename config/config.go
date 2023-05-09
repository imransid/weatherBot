package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	DatabaseURL string
}

var AppConfig *Config

// LoadConfig loads the application configuration
func LoadConfig() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to read configuration file")
	}

	AppConfig = &Config{
		DatabaseURL: "postgresql://admin:newpassword@127.0.0.1:5433/weatherApp",
		//viper.GetString(""),
	}
}

// NewPostgreSQLDB creates a new PostgreSQL database connection
func NewPostgreSQLDB() (*sql.DB, error) {
	connStr := "postgres://admin:newpassword@127.0.0.1:5433/weatherApp?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to the database successfully")

	return db, nil
}
