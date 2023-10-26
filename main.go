package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"stocks-crud-api/controllers"
	"stocks-crud-api/models"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database connection
	db, err := gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	))
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// AutoMigrate will create the table. This is for demonstration purposes.
	db.AutoMigrate(&models.Stock{})

	// Create a new instance of the StocksController and pass the database connection
	stocksController := controllers.NewStocksController(db)

	// Create a new Gin router
	r := gin.Default()

	// Define routes
	r.GET("/stocks", stocksController.GetAllStocks)
	r.GET("/stocks/:symbol", stocksController.GetStockBySymbol)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
