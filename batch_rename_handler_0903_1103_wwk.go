// 代码生成时间: 2025-09-03 11:03:54
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
)

// 文件重命名请求结构体
type RenameRequest struct {
    Source string `json:"source"` // 原文件名
    Target string `json:"target"` // 新文件名
}

// 文件重命名响应结构体
type RenameResponse struct {
    Success bool        `json:"success"` // 是否成功
    Message string     `json:"message"` // 消息
    Filename string    `json:"filename"` // 处理的文件名
}

// BatchRenameHandler 处理批量文件重命名
func BatchRenameHandler(c *gin.Context) {
    var requests []RenameRequest
    if err := c.ShouldBindJSON(&requests); err != nil {
        c.JSON(http.StatusBadRequest, RenameResponse{
            Success: false,
            Message: "Invalid request format",
        })
        return
    }

    for _, req := range requests {
        source := req.Source
        target := req.Target

        // 检查文件是否存在
        if _, err := os.Stat(source); os.IsNotExist(err) {
            c.JSON(http.StatusBadRequest, RenameResponse{
                Success: false,
                Message: fmt.Sprintf("File not found: %s", source),
                Filename: source,
            })
            return
        }

        // 检查是否是合法的文件路径
        if !filepath.ValidPath(target) {
            c.JSON(http.StatusBadRequest, RenameResponse{
                Success: false,
                Message: fmt.Sprintf("Invalid target path: %s", target),
                Filename: target,
            })
            return
        }

        // 重命名文件
        if err := os.Rename(source, target); err != nil {
            c.JSON(http.StatusInternalServerError, RenameResponse{
                Success: false,
                Message: fmt.Sprintf("Failed to rename file: %s, error: %s", source, err),
                Filename: source,
            })
            return
        }
    }

    c.JSON(http.StatusOK, RenameResponse{
        Success: true,
        Message: "Files renamed successfully",
    })
}

func main() {
    r := gin.Default()

    // 使用Gin中间件，例如Logger和Recovery
    r.Use(gin.Logger(), gin.Recovery())

    // 批量文件重命名的路由
    r.POST("/rename", BatchRenameHandler)

    // 启动服务
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %s", err)
    }
}
