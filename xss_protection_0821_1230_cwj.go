// 代码生成时间: 2025-08-21 12:30:51
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "html"
)

// XssMiddleware is a Gin middleware that provides basic XSS protection
// by sanitizing incoming request parameters to prevent cross-site scripting attacks.
func XssMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Sanitizing the query parameters
        c.Request.URL.RawQuery = html.EscapeString(c.Request.URL.RawQuery)
        // Sanitizing the form data
        if c.Request.Method == http.MethodPost {
            for key, values := range c.Request.PostForm {
                sanitizedValues := make([]string, len(values))
                for i, value := range values {
                    sanitizedValues[i] = html.EscapeString(value)
                }
                c.Request.PostForm[key] = sanitizedValues
            }
        }
        c.Next()
    }
}

// main function to start the Gin server
func main() {
    r := gin.Default()

    // Registering the XSS protection middleware
    r.Use(XssMiddleware())

    // Define a simple route that echoes back the input parameters
    r.GET("/", func(c *gin.Context) {
        // Get the sanitized input parameter
        userInput := c.DefaultQuery("input", "")

        // Respond with the sanitized input parameter
        c.JSON(http.StatusOK, gin.H{
            "input": userInput,
        })
    })

    // Start the server
    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
