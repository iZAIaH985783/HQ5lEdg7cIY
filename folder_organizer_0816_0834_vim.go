// 代码生成时间: 2025-08-16 08:34:20
package main

import (
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

// FolderOrganizerHandler 结构体，用于处理文件夹结构整理请求
type FolderOrganizerHandler struct {
    SourceDir string
}

// NewFolderOrganizerHandler 创建一个FolderOrganizerHandler实例
func NewFolderOrganizerHandler(sourceDir string) *FolderOrganizerHandler {
    return &FolderOrganizerHandler{
        SourceDir: sourceDir,
    }
}

// Organize 处理文件夹整理逻辑
func (f *FolderOrganizerHandler) Organize(c *gin.Context) {
    // 检查源目录是否存在
    if _, err := os.Stat(f.SourceDir); os.IsNotExist(err) {
        c.JSON(400, gin.H{
            "error": "source directory does not exist",
        })
        return
    }

    // 遍历目录并整理
    err := filepath.WalkDir(f.SourceDir, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if d.IsDir() {
            return nil
        }

        // 将文件移动到以文件名作为子目录的路径
        baseName := d.Name()
        newDir := filepath.Join(f.SourceDir, baseName)
        if err := os.MkdirAll(newDir, 0755); err != nil {
            return err
        }

        newFilePath := filepath.Join(newDir, baseName)
        if err := os.Rename(path, newFilePath); err != nil {
            return err
        }
        return nil
    })

    if err != nil {
        c.JSON(500, gin.H{
            "error": fmt.Sprintf("error organizing files: %v", err),
        })
        return
    }

    c.JSON(200, gin.H{
        "message": "files organized successfully",
    })
}

func main() {
    r := gin.Default()

    // 创建FolderOrganizerHandler实例
    handler := NewFolderOrganizerHandler("./myfolder")

    // 注册路由和处理器
    r.POST("/organize", handler.Organize)

    // 启动服务
    log.Fatal(r.Run(":8080"))
}
