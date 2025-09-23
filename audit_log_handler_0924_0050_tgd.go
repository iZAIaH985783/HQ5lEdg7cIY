// 代码生成时间: 2025-09-24 00:50:56
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// 日志记录中间件
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path
        method := c.Request.Method
        c.Next()

        // 记录请求的详细信息
        log.Printf("[AUDIT] %v %v %v %v", method, path, c.Request.RequestURI(), time.Since(start))
    }
}

// 错误处理中间件
func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        // 如果有错误发生，记录错误
        if len(c.Errors) > 0 {
            for _, e := range c.Errors {
                // 记录错误信息
                log.Printf("[ERROR] %v", e.Err)
            }
        }
    }
}

func main() {
    router := gin.Default()

    // 使用中间件
    router.Use(Logger(), ErrorHandler())

    // 测试路由
    router.GET("/test", func(c *gin.Context) {
        // 故意抛出一个错误
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Internal Server Error",
        })
    })

    // 启动服务器
    router.Run()
}
