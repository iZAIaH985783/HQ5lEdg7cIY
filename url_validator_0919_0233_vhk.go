// 代码生成时间: 2025-09-19 02:33:28
package main

import (
    "net/http"
    "net/url"
    "github.com/gin-gonic/gin"
    "strings"
)

// URLValidatorHandler Gin处理器用于验证URL链接的有效性
func URLValidatorHandler(c *gin.Context) {
    var requestBody struct {
        URL string `json:"url"`
    }

    // 解析请求体中的JSON数据
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid JSON format"
        })
        return
    }

    // 验证URL格式
    parsedURL, err := url.ParseRequestURI(requestBody.URL)
    if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid URL provided"
        })
        return
    }

    // 如果URL有效，返回成功消息
    c.JSON(http.StatusOK, gin.H{
        "message": "URL is valid"
    })
}

func main() {
    r := gin.Default()

    // 注册URL验证处理器
    r.POST("/validate-url", URLValidatorHandler)

    // 启动服务
    r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
