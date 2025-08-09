// 代码生成时间: 2025-08-09 08:02:35
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// CacheData is a struct to store cache data
type CacheData struct {
    Data []byte
    Expiry time.Time
}

// cacheStore is a simple in-memory cache store
var cacheStore = make(map[string]*CacheData)

// middlewareCacher is a middleware to handle caching
func middlewareCacher(c *gin.Context) {
    now := time.Now()
    key := c.Request.URL.Path
    if cacheData, exists := cacheStore[key]; exists && cacheData.Expiry.After(now) {
        c.Data(http.StatusOK, "application/json", cacheData.Data)
        c.Abort()
        return
    }
    c.Next()
    if cacheData, exists := cacheStore[key]; exists && cacheData.Expiry.After(now) {
        // Update expiry time
        cacheData.Expiry = now.Add(10 * time.Minute)
        cacheStore[key] = cacheData
    } else {
        // Cache the new data
        cacheData := &CacheData{
            Data:     c.Writer.Written,
            Expiry:   now.Add(10 * time.Minute),
        }
        cacheStore[key] = cacheData
    }
}

// cacheHandler is the handler that serves cached content or the original content
func cacheHandler(c *gin.Context) {
    // Simulating a slow process to demonstrate caching
    time.Sleep(5 * time.Second)
    c.JSON(http.StatusOK, gin.H{
        "message": "This is a cached response",
    })
    // After the response, we can store it in the cache
    cacheData := &CacheData{
        Data:     c.Writer.Written,
        Expiry:   time.Now().Add(10 * time.Minute),
    }
    cacheStore[c.Request.URL.Path] = cacheData
}

func main() {
    router := gin.Default()
    // Register the cacher middleware
    router.Use(middlewareCacher)

    // Register cacheHandler with a path
    router.GET("/cache", cacheHandler)

    // Start serving
    log.Fatal(router.Run(":8080"))
}
