// 代码生成时间: 2025-08-08 03:09:41
package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gin-gonic/gin"
)

// InventoryItem represents an item in inventory
type InventoryItem struct {
    ID    int    "json:"id""
    Name  string "json:"name""
    Count int    "json:"count""
}

var inventory = make(map[int]InventoryItem)
var nextID = 1

// Middleware to handle errors
func errorHandlingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        // Check if there is an error in the context
        if len(c.Errors) > 0 {
            // Handle the error
            err := c.Errors.Last().Err
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
        }
    }
}

// GetInventoryItem retrieves an inventory item by ID
func GetInventoryItem(c *gin.Context) {
    id := c.Param("id")
    item, exists := inventory[parseInt(id)]
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{
            "error": fmt.Sprintf("Inventory item with ID %s not found", id),
        })
        return
    }
    c.JSON(http.StatusOK, item)
}

// CreateInventoryItem adds a new inventory item
func CreateInventoryItem(c *gin.Context) {
    var newItem InventoryItem
    if err := c.ShouldBindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    inventory[nextID] = newItem
    newItem.ID = nextID
    nextID++
    c.JSON(http.StatusCreated, newItem)
}

// UpdateInventoryItem updates an existing inventory item
func UpdateInventoryItem(c *gin.Context) {
    var updatedItem InventoryItem
    if err := c.ShouldBindJSON(&updatedItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    id := c.Param("id")
    if _, exists := inventory[parseInt(id)]; !exists {
        c.JSON(http.StatusNotFound, gin.H{
            "error": fmt.Sprintf("Inventory item with ID %s not found", id),
        })
        return
    }
    inventory[parseInt(id)] = updatedItem
    c.JSON(http.StatusOK, updatedItem)
}

// DeleteInventoryItem deletes an inventory item by ID
func DeleteInventoryItem(c *gin.Context) {
    id := c.Param("id")
    if _, exists := inventory[parseInt(id)]; !exists {
        c.JSON(http.StatusNotFound, gin.H{
            "error": fmt.Sprintf("Inventory item with ID %s not found", id),
        })
        return
    }
    delete(inventory, parseInt(id))
    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("Inventory item with ID %s deleted", id),
    })
}

// parseInt is a helper function to convert string to int
func parseInt(id string) int {
    i, err := strconv.Atoi(id)
    if err != nil {
        log.Fatalf("Error converting ID to integer: %v", err)
        return 0
    }
    return i
}

func main() {
    router := gin.Default()

    // Register middleware for error handling
    router.Use(errorHandlingMiddleware())

    // Routes
    router.GET("/items/:id", GetInventoryItem)
    router.POST("/items", CreateInventoryItem)
    router.PUT("/items/:id", UpdateInventoryItem)
    router.DELETE("/items/:id", DeleteInventoryItem)

    // Start the server
    log.Fatal(router.Run(":8080"))
}
