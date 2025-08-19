// 代码生成时间: 2025-08-19 08:36:34
package main

import (
    "fmt"
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
)

// CacheKeyFn 定义了一个函数类型，用于根据请求生成缓存键
type CacheKeyFn func(c *gin.Context) string

// CacheData 定义了缓存数据的结构体
type CacheData struct {
    Data   []byte
    Expiry time.Time
}

// cacheMiddleware 是一个Gin中间件，用于实现缓存策略
func cacheMiddleware(cacheKeyFn CacheKeyFn, cache *map[string]CacheData) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 生成缓存键
        key := cacheKeyFn(c)
        // 检查缓存是否存在
        if cachedData, exists := (*cache)[key]; exists && !cachedData.Expiry.Before(time.Now()) {
            // 如果缓存未过期，则直接返回缓存数据
            c.Writer.Write(cachedData.Data)
            c.Abort()
            return
        }

        // 设置一个标记，用于后续操作完成后更新缓存
        c.Next()
        // 从响应中获取数据，并更新缓存
        if c.Writer.Status() == http.StatusOK {
            bodyData := c.Writer.Bytes()
            expiry := time.Now().Add(10 * time.Minute) // 缓存有效期为10分钟
            (*cache)[key] = CacheData{
                Data:   bodyData,
                Expiry: expiry,
            }
        }
    }
}

func main() {
    r := gin.Default()

    // 初始化缓存
    cache := make(map[string]CacheData)

    // 使用中间件，传入缓存键生成函数和缓存
    r.Use(cacheMiddleware(func(c *gin.Context) string {
        // 根据请求路径和查询参数生成缓存键
        return fmt.Sprintf("%s?%s", c.Request.URL.Path, c.Request.URL.RawQuery)
    }, &cache))

    // 定义一个简单的GET请求处理器
    r.GET("/test", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "This is a test response",
        })
    })

    // 启动服务器
    r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}
