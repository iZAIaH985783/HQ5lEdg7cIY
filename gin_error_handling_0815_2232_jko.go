// 代码生成时间: 2025-08-15 22:32:53
package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// ErrorResponse represents the structure of an error response in JSON format.
type ErrorResponse struct {
    Error string `json:"error"`
}

// ErrorHandler is a middleware that handles any panics that occur in downstream middlewares or handlers.
func ErrorHandler(c *gin.Context) {
    defer func() {
        if err := recover(); err != nil {
            c.Abort()
            resp := ErrorResponse{Error: fmt.Sprintf("%v", err)}
            c.JSON(http.StatusInternalServerError, resp)
        }
    }()
    c.Next()
}

// SimpleHandler is a simple Gin handler that can be used to demonstrate error handling.
func SimpleHandler(c *gin.Context) {
    // Simulate an error condition
    if err := fmt.Errorf("an error occurred"); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "everything is ok",
    })
}

func main() {
    router := gin.Default()

    // Register the middleware
    router.Use(ErrorHandler)

    // Define a route with the SimpleHandler
    router.GET("/", SimpleHandler)

    // Start the server
    router.Run(":8080")
}
