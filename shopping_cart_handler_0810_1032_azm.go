// 代码生成时间: 2025-08-10 10:32:50
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
)

// ShoppingCart represents the structure of a shopping cart
type ShoppingCart struct {
    Items map[string]int `json:"items"`
}

// NewShoppingCart initializes a new ShoppingCart with an empty map for items
func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{Items: make(map[string]int)}
}

// AddItemToCart adds an item to the shopping cart with a specified quantity
func (sc *ShoppingCart) AddItemToCart(item string, quantity int) {
    sc.Items[item] = quantity
}

// RemoveItemFromCart removes an item from the shopping cart
func (sc *ShoppingCart) RemoveItemFromCart(item string) {
    delete(sc.Items, item)
}

// GetCart returns the current state of the shopping cart
func (sc *ShoppingCart) GetCart() map[string]int {
    return sc.Items
}

// CartHandler handles cart-related requests
func CartHandler(c *gin.Context) {
    cart := NewShoppingCart()

    // Adding an item to the cart
    if c.Post("/add") {
        var addItem struct {
            Item    string `json:"item"`
            Quantity int    `json:"quantity"`
        }
        if err := c.ShouldBindJSON(&addItem); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid input"
            })
            return
        }
        cart.AddItemToCart(addItem.Item, addItem.Quantity)
        c.JSON(http.StatusOK, cart.GetCart())
    }

    // Removing an item from the cart
    if c.Delete("/remove") {
        var removeItem struct {
            Item string `json:"item"`
        }
        if err := c.ShouldBindJSON(&removeItem); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid input"
            })
            return
        }
        cart.RemoveItemFromCart(removeItem.Item)
        c.JSON(http.StatusOK, cart.GetCart())
    }

    // Getting the current cart
    if c.Get("/get") {
        c.JSON(http.StatusOK, cart.GetCart())
    }
}

func main() {
    // Initialize Gin router
    router := gin.Default()

    // Register the cart handler with Gin
    router.POST("/cart/add", CartHandler)
    router.DELETE("/cart/remove", CartHandler)
    router.GET("/cart/get", CartHandler)

    // Start the server
    router.Run()
}
