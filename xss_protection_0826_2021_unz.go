// 代码生成时间: 2025-08-26 20:21:10
package main

import (
    "github.com/gin-gonic/gin"
    "html"
    "net/http"
    "strings"
)

// XSSProtectionMiddleware 是一个中间件，用于防止XSS攻击。
// 它通过HTML转义所有输入来实现。
func XSSProtectionMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        for _, value := range c.Request.URL.Query() {
            safeValue := html.EscapeString(value[0])
            c.Request.URL.RawQuery = strings.ReplaceAll(c.Request.URL.RawQuery, value[0], safeValue)
        }

        for key, value := range c.Request.PostForm {
            safeValue := html.EscapeString(value[0])
            c.Request.PostForm[key] = []string{safeValue}
        }

        c.Next()
    }
}

func main() {
    // 创建一个新的Gin路由器
    router := gin.Default()

    // 使用XSS防护中间件
    router.Use(XSSProtectionMiddleware())

    // 定义一个简单的处理函数，用于测试
    router.GET("/test", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World!",
        })
    })

    // 启动服务器
    router.Run(":8080")
}
