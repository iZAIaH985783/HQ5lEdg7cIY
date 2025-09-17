// 代码生成时间: 2025-09-17 21:46:08
package main

import (
    "fmt"
    "math"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// ErrorResponse defines the structure for error responses.
type ErrorResponse struct {
    Error string `json:"error"`
}

// MathService handles mathematical operations.
func MathService(c *gin.Context) {
    // Retrieve operation type and numbers from the URL parameters.
    operation := c.Param("operation")
    num1, err1 := strconv.ParseFloat(c.Query("num1"), 64)
    num2, err2 := strconv.ParseFloat(c.Query("num2"), 64)

    // Check for errors in parameter parsing.
    if err1 != nil || err2 != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input parameters"})
        return
    }

    // Perform the requested mathematical operation.
    switch operation {
    case "add":
        result := num1 + num2
        c.JSON(http.StatusOK, gin.H{"result": result})
    case "subtract":
        result := num1 - num2
        c.JSON(http.StatusOK, gin.H{"result": result})
    case "multiply":
        result := num1 * num2
        c.JSON(http.StatusOK, gin.H{"result": result})
    case "divide":
        if num2 == 0 {
            c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Cannot divide by zero"})
            return
        }
        result := num1 / num2
        c.JSON(http.StatusOK, gin.H{"result": result})
    default:
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Unsupported operation"})
    }
}

func main() {
    r := gin.Default()

    // Define routes for the math operations.
    r.GET("/add/:num1/:num2", MathService)
    r.GET("/subtract/:num1/:num2", MathService)
    r.GET("/multiply/:num1/:num2", MathService)
    r.GET("/divide/:num1/:num2", MathService)

    // Start the server.
    r.Run() // listening and serving on 0.0.0.0:8080
}
