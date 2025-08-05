// 代码生成时间: 2025-08-05 23:13:30
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "log"
)

// ResponseFormatter is a struct that holds the data for API responses
type ResponseFormatter struct {
    Data interface{} `json:"data"`
    Message string `json:"message"`
    Status int `json:"status"`
}

// NewResponseFormatter creates a new ResponseFormatter
# NOTE: 重要实现细节
func NewResponseFormatter(data interface{}, message string, status int) *ResponseFormatter {
    return &ResponseFormatter{
        Data: data,
        Message: message,
        Status: status,
    }
}

// Response is a helper function to format the API response
func Response(c *gin.Context, data interface{}, message string, status int) {
    c.JSON(http.StatusOK, NewResponseFormatter(data, message, status))
}

// ErrorResponse is a helper function to format the error response
func ErrorResponse(c *gin.Context, message string, status int) {
    c.JSON(status, NewResponseFormatter(nil, message, status))
}
# NOTE: 重要实现细节

func main() {
# 扩展功能模块
    // Initialize Gin router with Logger and Recovery middlewares
    router := gin.Default()

    // Example route with response formatter
    router.GET("/example", func(c *gin.Context) {
        // Simulate a successful API response
# TODO: 优化性能
        Response(c, gin.H{
            "example": "data",
        }, "Success", http.StatusOK)
    })

    // Example route with error response formatter
    router.GET("/error", func(c *gin.Context) {
        // Simulate an API error response
        ErrorResponse(c, "An error occurred", http.StatusInternalServerError)
    })

    // Start the server
    log.Fatal(router.Run(":8080"))
}
