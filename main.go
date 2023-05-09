// // package main

// // import (
// // 	"net/http"

// // 	"github.com/gin-gonic/gin"
// // )

// // // Struct to represent weather data
// // type WeatherData struct {
// // 	Temperature string
// // 	Conditions  string
// // }

// // func main() {
// // 	router := gin.Default()

// // 	// Define a route for the weather endpoint
// // 	router.GET("/weather/:city", func(c *gin.Context) {
// // 		city := c.Param("city")

// // 		// Call a function to fetch weather data for the specified city
// // 		weatherData := getWeatherData(city)

// // 		// Return the weather data as JSON
// // 		c.JSON(http.StatusOK, gin.H{
// // 			"city":       city,
// // 			"temperature": weatherData.Temperature,
// // 			"conditions": weatherData.Conditions,
// // 		})
// // 	})

// // 	// Run the server
// // 	router.Run(":8080")
// // }

// // // Function to fetch weather data from an external API
// // func getWeatherData(city string) WeatherData {
// // 	// Here you can implement the logic to fetch weather data from an
// // //external API
// // 	// For this example, we'll return some dummy data
// // 	return WeatherData{
// // 		Temperature: "25°C",
// // 		Conditions:  "Sunny",
// // 	}
// // }

package main

import (
	"log"

	"example.com/weatherApp/api"
	"example.com/weatherApp/config"
	"example.com/weatherApp/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load the application configuration
	config.LoadConfig()

	// Initialize the Gin router
	router := gin.Default()

	// Create a new PostgreSQL connection
	db, err := config.NewPostgreSQLDB()
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	// Initialize the weather repository
	weatherRepo := repositories.NewWeatherRepository(db)

	// Initialize the API routes
	api.InitRoutes(router, weatherRepo)

	// Start the server
	router.Run(":8080")
}

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	connStr := "postgres://admin:newpassword@127.0.0.1:5433/weatherApp?sslmode=disable"

// 	db, err := sql.Open("postgres", connStr)

// 	if err != nil {
// 		log.Fatalf("failed to open database connection: %v", err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Errorf("failed to ping database: %w", err)
// 	}

// 	log.Println("Connected to the database successfully")

// 	// return db, nil
// }

// // NewPostgreSQLDB creates a new PostgreSQL database connection
// // func NewPostgreSQLDB() (*sql.DB, error) {
// //     connStr := "postgres://admin:newpassword@127.0.0.1:5433/weatherApp?sslmode=disable"

// //     db, err := sql.Open("postgres", connStr)

// //     if err != nil {
// //         log.Fatalf("failed to open database connection: %v", err)
// //     }

// //     err = db.Ping()
// //     if err != nil {
// //         return nil, fmt.Errorf("failed to ping database: %w", err)
// //     }

// //     log.Println("Connected to the database successfully")

// //     return db, nil
// // }
