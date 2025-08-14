// 代码生成时间: 2025-08-14 18:05:18
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
)

// HashCalculatorHandler 结构体用于接收请求并计算哈希值
type HashCalculatorHandler struct {}

// CalculateSha256 计算SHA-256哈希值
func (h *HashCalculatorHandler) CalculateSha256(c *gin.Context) {
    // 从请求中获取数据
    input := c.PostForm("input")
    if input == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Input is required",
        })
        return
    }

    // 计算SHA-256哈希值
    hash := sha256.Sum256([]byte(input))
    result := hex.EncodeToString(hash[:])

    // 返回结果
    c.JSON(http.StatusOK, gin.H{
        "hash": result,
    })
}

// SetupRouter 设置Gin路由器
func SetupRouter() *gin.Engine {
    r := gin.Default()
    // 创建HashCalculatorHandler实例
    handler := &HashCalculatorHandler{}

    // 将处理器绑定到路由上
    r.POST("/hash", handler.CalculateSha256)

    return r
}

func main() {
    // 设置路由
    router := SetupRouter()

    // 启动服务
    log.Fatal(router.Run(":8080"))
}
