// 代码生成时间: 2025-08-18 19:59:01
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
)

// FolderStructureHandler 结构体用于处理文件夹结构整理的请求
type FolderStructureHandler struct {
    // RootPath 是需要整理的文件夹的根路径
    RootPath string
}

// NewFolderStructureHandler 创建并返回一个FolderStructureHandler实例
func NewFolderStructureHandler(rootPath string) *FolderStructureHandler {
    return &FolderStructureHandler{
        RootPath: rootPath,
    }
}

// HandleFolderStructure 处理POST请求，整理文件夹结构
func (h *FolderStructureHandler) HandleFolderStructure(c *gin.Context) {
    // 检查根路径是否存在
    if _, err := os.Stat(h.RootPath); os.IsNotExist(err) {
        c.JSON(404, gin.H{
            "error": "Root path does not exist",
        })
        return
    }

    files, err := os.ReadDir(h.RootPath)
    if err != nil {
        c.JSON(500, gin.H{
            "error": fmt.Sprintf("Failed to read directory: %v", err),
        })
        return
    }

    var folderStructure []string
    for _, file := range files {
        if file.IsDir() {
            folderStructure = append(folderStructure, file.Name())
        }
    }

    // 将整理后的文件夹结构返回给客户端
    c.JSON(200, gin.H{
        "folder_structure": folderStructure,
    })
}

func main() {
    // 设置Gin的模式为Release模式以提高性能
    gin.SetMode(gin.ReleaseMode())

    // 初始化Gin引擎
    r := gin.Default()

    // 设置根路径为当前工作目录，可以根据自己的需要修改
    rootPath := "."

    // 创建FolderStructureHandler实例
    handler := NewFolderStructureHandler(rootPath)

    // 定义处理POST请求的路由
    r.POST("/folder_structure", handler.HandleFolderStructure)

    // 启动服务器
    r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
