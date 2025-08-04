// 代码生成时间: 2025-08-05 00:26:26
package main

import (
    "encoding/json"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

// ShoppingCart represents the shopping cart structure.
type ShoppingCart struct {
    Items map[string]int `json:"items"`
}

// CartItem represents an item in the shopping cart.
type CartItem struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Price float64 `json:"price"`
}

// NewShoppingCart creates a new shopping cart.
func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{Items: make(map[string]int)}
}

// AddItem adds an item to the shopping cart.
func (cart *ShoppingCart) AddItem(item CartItem) {
    cart.Items[item.ID] += 1
}

// RemoveItem removes an item from the shopping cart.
func (cart *ShoppingCart) RemoveItem(itemID string) {
    if _, exists := cart.Items[itemID]; exists {
        delete(cart.Items, itemID)
    }
}

func main() {
    r := gin.Default()

    // Middleware to handle JSON responses.
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Content-Type", "application/json")
    })

    // Initialize shopping cart.
    cart := NewShoppingCart()

    // Add item to cart.
    r.POST("/add", func(c *gin.Context) {
        var item CartItem
        if err := c.ShouldBindJSON(&item); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }
        cart.AddItem(item)
        c.JSON(http.StatusOK, cart.Items)
    })

    // Remove item from cart.
    r.POST("/remove", func(c *gin.Context) {
        var itemID string
        if err := c.ShouldBindJSON(&itemID); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }
        cart.RemoveItem(itemID)
        c.JSON(http.StatusOK, cart.Items)
    })

    // Get cart items.
    r.GET("/items", func(c *gin.Context) {
        c.JSON(http.StatusOK, cart.Items)
    })

    // Start the server.
    r.Run()
}
