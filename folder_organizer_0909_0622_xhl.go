// 代码生成时间: 2025-09-09 06:22:34
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
)

// FolderOrganizerHandler 结构用于处理文件夹整理请求
type FolderOrganizerHandler struct {
    RootPath string
}

// NewFolderOrganizerHandler 创建一个新的 FolderOrganizerHandler 实例
func NewFolderOrganizerHandler(rootPath string) *FolderOrganizerHandler {
    return &FolderOrganizerHandler{
        RootPath: rootPath,
    }
}

// Organize 是处理文件夹整理的函数
// 它将遍历指定的根路径，按文件类型组织文件到子文件夹中
func (f *FolderOrganizerHandler) Organize(c *gin.Context) {
    var err error
    defer func() {
        if err != nil {
            c.JSON(500, gin.H{
                "error": err.Error(),
            })
        }
    }()

    // 检查根路径是否存在
    if _, err := os.Stat(f.RootPath); os.IsNotExist(err) {
        err = fmt.Errorf("root path does not exist: %s", f.RootPath)
        return
    }

    // 遍历根路径下的所有文件和文件夹
    err = filepath.WalkDir(f.RootPath, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }

        // 忽略目录
        if d.IsDir() {
            return nil
        }

        // 根据文件扩展名组织文件到子文件夹
        fileExt := strings.TrimPrefix(strings.TrimPrefix(filepath.Ext(path), "."), ".")
        if len(fileExt) == 0 {
            return nil // 忽略无扩展名的文件
        }

        destDir := filepath.Join(f.RootPath, fileExt)
        if _, err := os.Stat(destDir); os.IsNotExist(err) {
            err = os.MkdirAll(destDir, 0755)
            if err != nil {
                return err
            }
        }

        // 移动文件到对应的子文件夹
        destPath := filepath.Join(destDir, filepath.Base(path))
        err = os.Rename(path, destPath)
        if err != nil {
            return err
        }

        return nil
    })

    if err != nil {
        c.JSON(200, gin.H{
            "message": "Files organized successfully",
        })
    }
}

func main() {
    r := gin.Default()

    // 使用中间件记录请求日志
    r.Use(gin.Logger())

    // 使用中间件恢复任何panic，返回500错误响应
    r.Use(gin.Recovery())

    // 设置文件夹整理处理器
    organizer := NewFolderOrganizerHandler("/path/to/your/folder")
    r.POST("/organize", organizer.Organize)

    // 启动服务
    r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
