package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"stocks-crud-api/controllers"
)

// InitializeRoutes initializes all routes and sets up the Gin router.
func InitializeRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Create an instance of the StocksController
	stocksController := controllers.NewStocksController(db)

	// Define routes
	api := router.Group("/api")
	{
		api.GET("/stocks", stocksController.GetAllStocks)
		api.GET("/stocks/:symbol", stocksController.GetStockBySymbol)
	}

	return router
}
