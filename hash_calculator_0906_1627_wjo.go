// 代码生成时间: 2025-09-06 16:27:01
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
)

// HashCalculatorHandler 处理哈希值计算请求
func HashCalculatorHandler(c *gin.Context) {
    input := c.PostForm("input") // 从请求中获取输入值
    if input == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Input is required",
        })
        return
    }

    // 使用SHA-256算法计算哈希值
    hash := sha256.Sum256([]byte(input))
    hashString := hex.EncodeToString(hash[:])

    // 返回哈希值
    c.JSON(http.StatusOK, gin.H{
        "hash": hashString,
    })
}

func main() {
    r := gin.Default() // 创建Gin路由器

    // 注册哈希值计算处理器
    r.POST("/hash", HashCalculatorHandler)

    // 启动服务器
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

// 请注意，此代码不包括单元测试，实际项目中应该包含适当的测试来验证功能。