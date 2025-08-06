// 代码生成时间: 2025-08-06 12:11:56
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// ErrorLogMiddleware is a Gin middleware that logs errors.
func ErrorLogMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Capture the start time to calculate latency later.
        startTime := time.Now()

        // Continue to the next middleware or handler
        c.Next()

        // Check if there's an error using the error key in the context.
        if len(c.Errors.ByType(gin.ErrorTypeAny)) > 0 {
            // Retrieve the error message.
            errorMsg := c.Errors.Last().Err.Error()

            // Log the error with request information.
            log.Printf("[ERROR] - %s - %s %s %s %s", errorMsg, c.Request.Method, c.Request.URL.Path,
                c.ClientIP(), time.Since(startTime))
        }
    }
}

func main() {
    r := gin.New()

    // Use the error logging middleware.
    r.Use(ErrorLogMiddleware())

    // Example route that intentionally causes an error.
    r.GET("/error", func(c *gin.Context) {
        c.String(http.StatusInternalServerError, "Internal Server Error")
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "error": "Something went wrong.",
        })
    })

    // Start the server.
    r.Run() // listening on 0.0.0.0:8080
}
