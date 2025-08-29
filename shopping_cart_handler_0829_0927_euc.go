// 代码生成时间: 2025-08-29 09:27:13
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "encoding/json"
)

// ShoppingCart represents a shopping cart with items and their quantities
type ShoppingCart struct {
    Items map[string]int `json:"items"`
}

// NewShoppingCart initializes a new shopping cart
func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{Items: make(map[string]int)}
}

// AddItem adds an item to the cart with a specified quantity
func (cart *ShoppingCart) AddItem(item string, quantity int) {
    if _, exists := cart.Items[item]; exists {
        cart.Items[item] += quantity
    } else {
        cart.Items[item] = quantity
    }
}

// RemoveItem removes an item from the cart
func (cart *ShoppingCart) RemoveItem(item string) {
    if _, exists := cart.Items[item]; exists {
        delete(cart.Items, item)
    }
}

// CartHandler handles HTTP requests for the shopping cart
func CartHandler(cart *ShoppingCart) gin.HandlerFunc {
    return func(c *gin.Context) {
        switch c.Request.Method {
        case http.MethodGet:
            // Return the current state of the cart
            c.JSON(http.StatusOK, cart)
        case http.MethodPost:
            // Add an item to the cart
            var item struct {
                Item   string `json:"item"`
                Quantity int    `json:"quantity"`
            }
            if err := c.ShouldBindJSON(&item); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{'error': 'Invalid input'})
                return
            }
            cart.AddItem(item.Item, item.Quantity)
            c.JSON(http.StatusOK, cart)
        case http.MethodDelete:
            // Remove an item from the cart
            var item struct {
                Item string `json:"item"`
            }
            if err := c.ShouldBindJSON(&item); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{'error': 'Invalid input'})
                return
            }
            cart.RemoveItem(item.Item)
            c.JSON(http.StatusOK, cart)
        default:
            c.JSON(http.StatusMethodNotAllowed, gin.H{'error': 'Method not allowed'})
        }
    }
}

func main() {
    router := gin.Default()

    // Initialize a new shopping cart
    cart := NewShoppingCart()

    // Register the cart handler with all supported methods
    router.Handle(http.MethodGet, http.MethodPost, http.MethodDelete, "/cart", CartHandler(cart))

    // Start the server
    router.Run(":8080") // listening and serving on 0.0.0.0:8080
}
