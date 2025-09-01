// 代码生成时间: 2025-09-01 17:03:03
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "time"

    "github.com/gin-gonic/gin"
)

// FileSyncHandler 处理文件备份和同步的请求
func FileSyncHandler(c *gin.Context) {
    srcPath := c.DefaultQuery("src", "")
    destPath := c.DefaultQuery("dest", "")
    if srcPath == "" || destPath == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Source and destination paths are required",
        })
        return
    }

    // 创建备份的目录
    if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to create destination directory: %v", err),
        })
        return
    }

    // 获取源目录中的文件列表
    files, err := os.ReadDir(srcPath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to read source directory: %v", err),
        })
        return
    }

    for _, file := range files {
        srcFile, _ := filepath.Abs(filepath.Join(srcPath, file.Name()))
        destFile, _ := filepath.Abs(filepath.Join(destPath, file.Name()))

        // 检查文件是否已经存在于目标目录
        if _, err := os.Stat(destFile); os.IsNotExist(err) {
            // 复制文件
            if err := copyFile(srcFile, destFile); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error": fmt.Sprintf("Failed to copy file %s to %s: %v", srcFile, destFile, err),
                })
                return
            }
        } else {
            // 文件已存在，检查是否需要更新
            if needUpdate(srcFile, destFile) {
                if err := copyFile(srcFile, destFile); err != nil {
                    c.JSON(http.StatusInternalServerError, gin.H{
                        "error": fmt.Sprintf("Failed to update file %s to %s: %v", srcFile, destFile, err),
                    })
                    return
                }
            }
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Files have been backed up and synchronized successfully",
    })
}

// copyFile 复制文件
func copyFile(src, dest string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()

    destFile, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer destFile.Close()

    _, err = destFile.ReadFrom(sourceFile)
    return err
}

// needUpdate 检查是否需要更新文件
func needUpdate(src, dest string) bool {
    srcInfo, err := os.Stat(src)
    if err != nil {
        log.Fatal(err)
    }
    destInfo, err := os.Stat(dest)
    if err != nil {
        log.Fatal(err)
    }
    return srcInfo.ModTime().After(destInfo.ModTime())
}

func main() {
    r := gin.Default()
    r.GET("/sync", FileSyncHandler)
    r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}
