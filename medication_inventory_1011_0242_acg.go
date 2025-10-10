// 代码生成时间: 2025-10-11 02:42:24
package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// Medication represents a drug with its name and quantity
type Medication struct {
    Name    string `json:"name"`
    Quantity int    `json:"quantity"`
}

// Inventory represents a list of medications
type Inventory struct {
    Medications map[string]Medication
}

// NewInventory creates a new inventory with an empty map of medications
func NewInventory() *Inventory {
    return &Inventory{Medications: make(map[string]Medication)}
}

func main() {
    // Create a router without any middleware by default
    r := gin.Default()

    // Create a new inventory
    inventory := NewInventory()

    // Register a new medication handler
    r.POST("/add_medication", func(c *gin.Context) {
        // Bind the JSON to a Medication struct
        var m Medication
        if err := c.ShouldBindJSON(&m); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
            return
        }
        // Add the medication to the inventory
        inventory.Medications[m.Name] = m
        c.JSON(http.StatusOK, gin.H{"message": "Medication added successfully"})
    })

    // Register a retrieve medication handler
    r.GET("/medication/:name", func(c *gin.Context) {
        name := c.Param("name")
        m, exists := inventory.Medications[name]
        if !exists {
            c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Medication %s not found", name)})
            return
        }
        c.JSON(http.StatusOK, m)
    })

    // Register a update medication handler
    r.PUT("/update_medication/:name", func(c *gin.Context) {
        name := c.Param("name\)
        var m Medication
        if err := c.ShouldBindJSON(&m); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
            return
        }
        // Update the medication in the inventory
        if _, exists := inventory.Medications[name]; exists {
            inventory.Medications[name] = m
            c.JSON(http.StatusOK, gin.H{"message": "Medication updated successfully"})
        } else {
            c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Medication %s not found", name)})
        }
    })

    // Start the server
    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
