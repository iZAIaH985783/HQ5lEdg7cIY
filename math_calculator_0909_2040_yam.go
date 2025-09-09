// 代码生成时间: 2025-09-09 20:40:07
package main

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// MathCalculator defines the structure for the math calculator handler.
type MathCalculator struct {
}

// Handler for the math calculator routes.
func (m *MathCalculator) Handler() gin.HandlerFunc {
    return func(c *gin.Context) {
        var result float64
        var err error

        // Define the operations.
        operations := map[string]func(a, b float64) (float64, error){
            "add": func(a, b float64) (float64, error) {
                return a + b, nil
            },
            "subtract": func(a, b float64) (float64, error) {
                return a - b, nil
            },
            "multiply": func(a, b float64) (float64, error) {
                return a * b, nil
            },
            "divide": func(a, b float64) (float64, error) {
                if b == 0 {
                    return 0, fmt.Errorf("cannot divide by zero")
                }
                return a / b, nil
            },
        }

        // Get the operation from the query parameter.
        operation := c.Query("operation")
        if operation == "" {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "operation parameter is missing"
            })
            return
        }

        // Check if the operation is supported.
        if _, exists := operations[operation]; !exists {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "unsupported operation"
            })
            return
        }

        // Get the numbers from the query parameters.
        num1Str := c.Query("number1")
        num2Str := c.Query("number2")
        if num1Str == "" || num2Str == "" {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "number1 and number2 parameters are required"
            })
            return
        }

        num1, err := strconv.ParseFloat(num1Str, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "invalid number1 format"
            })
            return
        }
        num2, err := strconv.ParseFloat(num2Str, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "invalid number2 format"
            })
            return
        }

        // Perform the operation.
        result, err = operations[operation](num1, num2)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error()
            })
            return
        }

        // Return the result of the operation.
        c.JSON(http.StatusOK, gin.H{
            "result": result,
        })
    }
}

func main() {
    router := gin.Default()
    mathCalculator := MathCalculator{}

    // Register the math calculator route.
    router.GET("/math", mathCalculator.Handler())

    // Start the server.
    router.Run(":8080") // listening and serving on 0.0.0.0:8080
}