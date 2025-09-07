// 代码生成时间: 2025-09-08 01:54:38
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// ApiResponse represents the structured response format
type ApiResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data"`
    Message string     `json:"message"`
    Error   string     `json:"error"`
}

// Response is a helper function to format response data
func Response(c *gin.Context, data interface{}, message string, err error) {
    apiResponse := ApiResponse{
        Success: err == nil,
        Data:    data,
        Message: message,
        Error:   "",
    }

    if err != nil {
        apiResponse.Error = err.Error()
    }

    c.JSON(http.StatusOK, apiResponse)
}

// HandleError is a helper function to handle errors and return a formatted error response
func HandleError(c *gin.Context, err error, message string, statusCode int) {
    c.JSON(statusCode, ApiResponse{
        Success: false,
        Data:    nil,
        Message: message,
        Error:   err.Error(),
    })
}

// NewGinEngine creates a new Gin engine with necessary middlewares
func NewGinEngine() *gin.Engine {
    r := gin.Default()
    // You can add additional middlewares if needed
    // r.Use(gin.Recovery())
    // r.Use(gin.Logger())
    return r
}

// main function to start the server
func main() {
    r := NewGinEngine()

    // Define a sample route using the API response formatter
    r.GET("/example", func(c *gin.Context) {
        // Simulate successful response
        Response(c, gin.H{"example": "Hello World!"}, "Success", nil)
    })

    // Define another sample route with an error
    r.GET("/error", func(c *gin.Context) {
        // Simulate an error
        HandleError(c, ErrExample{}, "Something went wrong", http.StatusInternalServerError)
    })

    // Start the server
    r.Run()
}

// ErrExample is a mock error type for demonstration purposes
type ErrExample struct{}

func (e ErrExample) Error() string {
    return "Example error occurred"
}