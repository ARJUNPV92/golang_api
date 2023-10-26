package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "stocks-crud-api/models"
)

// StocksController is responsible for handling CRUD operations on stocks.
type StocksController struct {
    DB *gorm.DB
}

// NewStocksController creates a new instance of the StocksController.
func NewStocksController(db *gorm.DB) *StocksController {
    return &StocksController{DB: db}
}

// GetAllStocks returns a list of all stocks.
func (c *StocksController) GetAllStocks(ctx *gin.Context) {
    var stocks []models.Stock

    if err := c.DB.Find(&stocks).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stocks"})
        return
    }

    ctx.JSON(http.StatusOK, stocks)
}

// GetStockBySymbol retrieves a stock by its SYMBOL.
func (c *StocksController) GetStockBySymbol(ctx *gin.Context) {
    symbol := ctx.Param("symbol")
    var stock models.Stock

    if err := c.DB.Where("SYMBOL = ?", symbol).First(&stock).Error; err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
        return
    }

    ctx.JSON(http.StatusOK, stock)
}
