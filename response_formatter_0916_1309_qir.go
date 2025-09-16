// 代码生成时间: 2025-09-16 13:09:34
package main

import (
    "net/http"
    "log"
    "github.com/gin-gonic/gin"
)

// ApiResponse 结构体用于格式化API响应
type ApiResponse struct {
    Code    int         `json:"code"`    // 响应代码
    Message string      `json:"message"`  // 响应消息
    Data    interface{} `json:"data"`    // 响应数据
}

// NewApiResponse 初始化ApiResponse
func NewApiResponse(code int, message string, data interface{}) ApiResponse {
    return ApiResponse{
        Code:    code,
        Message: message,
        Data:    data,
    }
}

// APIResponseMiddleware Gin中间件，用于格式化API响应
func APIResponseMiddleware(c *gin.Context) {
    c.Next()

    // 获取状态码
    status := c.Writer.Status()
    // 获取响应数据
    response := make(map[string]interface{})
    response["code"] = status
    response["message"] = http.StatusText(status) // 根据状态码获取状态文本
    // 将响应写入客户端
    c.JSON(status, response)
}

// ErrorHandler Gin中间件，用于错误处理
func ErrorHandler(c *gin.Context) {
    c.Next()
    if len(c.Errors) > 0 {
        for _, e := range c.Errors {
            // 错误处理
            log.Printf("Error: %v
", e)
            c.JSON(http.StatusInternalServerError, gin.H{
                "code":    http.StatusInternalServerError,
                "message": e.Err,
            })
        }
    }
}

func main() {
    r := gin.Default()

    // 使用中间件
    r.Use(APIResponseMiddleware, ErrorHandler)

    // 测试API
    r.GET("/test", func(c *gin.Context) {
        // 正常响应
        c.JSON(http.StatusOK, NewApiResponse(http.StatusOK, "OK", gin.H{
            "hello": "world",
        }))
    })

    // 启动服务
    r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}
