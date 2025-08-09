// 代码生成时间: 2025-08-09 16:30:30
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// ErrorLoggerMiddleware is a middleware function that logs errors
func ErrorLoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        // Check if there is an error in the context
        if len(c.Errors.ByType(gin.ErrorTypeAny)) > 0 {
            // Retrieve the error
            err := c.Errors.Last().Err
            // Log the error with a timestamp
            log.Printf("[ERROR] - %s - %v", time.Now().Format(time.RFC3339), err)
            // Optionally, you can also add logging for the request details
            log.Printf("[ERROR] - Request Method: %s, Path: %s", c.Request.Method, c.Request.URL.Path)
        }
    }
}

// ErrorHandler is a function to handle errors
func ErrorHandler(c *gin.Context) {
    // Retrieve the error from the context
    err := c.Errors.Last().Err
    // Log the error
    log.Printf("[ERROR] - %s - %v", time.Now().Format(time.RFC3339), err)
    // Optionally, log the request details
    log.Printf("[ERROR] - Request Method: %s, Path: %s", c.Request.Method, c.Request.URL.Path)
    // Respond with an error message and HTTP status code 500
    c.JSON(http.StatusInternalServerError, gin.H{
        "error": "Internal Server Error",
    })
}

func main() {
    r := gin.Default()

    // Register the error logger middleware
    r.Use(ErrorLoggerMiddleware())

    // Register a sample route with an error handler
    r.GET("/error", func(c *gin.Context) {
        // Intentionally cause an error
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "error": "Error occurred"},
        )
    })

    // Use the default error handler from Gin, but you could also use your own custom handler
    r.NoRoute(ErrorHandler)
    r.NoMethod(ErrorHandler)

    // Start the server
    r.Run(":8080")
}
