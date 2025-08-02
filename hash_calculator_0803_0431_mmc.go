// 代码生成时间: 2025-08-03 04:31:53
 * Provides a simple HTTP endpoint to calculate and return the hash
 * of a given string.
 */

package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

// CalculateHash calculates the SHA-256 hash of the given string.
func CalculateHash(input string) (string, error) {
    hash := sha256.Sum256([]byte(input))
    return hex.EncodeToString(hash[:]), nil
}

// HashHandler is a Gin handler function that calculates and returns the hash of the input string.
func HashHandler(c *gin.Context) {
    input := c.PostForm("input")
    if input == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Missing input parameter.",
        })
        return
    }
    hash, err := CalculateHash(input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to calculate hash.",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "hash": hash,
    })
}

func main() {
    router := gin.Default()

    // Define the endpoint for the hash calculation.
    router.POST("/hash", HashHandler)

    // Run the server on port 8080.
    router.Run(":8080")
}