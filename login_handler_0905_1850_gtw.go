// 代码生成时间: 2025-09-05 18:50:41
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "log"
)

// LoginData represents the structure for login data
type LoginData struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func main() {
    r := gin.Default()

    // Define a route for user login with POST method
    r.POST("/login", loginHandler)

    // Start the server on port 8080
    r.Run(":8080")
}

// loginHandler handles the login request
func loginHandler(c *gin.Context) {
    var login LoginData
    if err := c.ShouldBindJSON(&login); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
        return
    }

    // Here you would normally validate the credentials against a database or authentication service
    // For demonstration purposes, we are using hardcoded credentials
    if login.Username != "admin" || login.Password != "password" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // If credentials are correct, return a success message
    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
