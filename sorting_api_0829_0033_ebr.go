// 代码生成时间: 2025-08-29 00:33:20
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// SortAlgorithmHandler is the handler for sorting algorithms
func SortAlgorithmHandler(c *gin.Context) {
    var numbers []int
    // BindJSON binds the JSON to a destination (pointer)
    if err := c.BindJSON(&numbers); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid input: must be an array of integers",
        })
        return
    }

    // Sorting the slice of integers using a simple bubble sort algorithm
    for i := 0; i < len(numbers); i++ {
        for j := 0; j < len(numbers)-i-1; j++ {
            if numbers[j] > numbers[j+1] {
                // Swap the elements
                numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
            }
        }
    }

    // Return the sorted slice as JSON response
    c.JSON(http.StatusOK, gin.H{
        "sorted_numbers": numbers,
    })
}

func main() {
    r := gin.Default()

    // Registering the handler for sorting
    r.POST("/sort", SortAlgorithmHandler)

    // Start serving on port 8080
    r.Run(":8080")
}
