package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"stocks-crud-api/models"
)

func TestGetAllStocks(t *testing.T) {
	// Setup a test Gin router
	router := gin.New()

	// Use a test database or create a mock for gorm.DB
	// ...

	// Create a new instance of the StocksController with the test DB
	stocksController := NewStocksController(&gorm.DB{})

	// Define the route
	router.GET("/api/stocks", stocksController.GetAllStocks)

	// Create a GET request to /api/stocks
	req, err := http.NewRequest("GET", "/api/stocks", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP recorder
	recorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var stocks []models.Stock
	if err := json.Unmarshal(recorder.Body.Bytes(), &stocks); err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	// Add assertions on the stocks or other aspects of the response
	// For example, you might check if the returned data is not empty
	if len(stocks) == 0 {
		t.Error("Expected non-empty list of stocks, got empty list")
	}

	// Add more assertions based on your specific use case
	// ...

}
