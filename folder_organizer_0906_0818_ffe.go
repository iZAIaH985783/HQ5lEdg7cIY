// 代码生成时间: 2025-09-06 08:18:32
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
)

// FolderOrganizer 结构体，用于存储文件夹路径
type FolderOrganizer struct {
    Path string
}

// NewFolderOrganizer 创建一个新的 FolderOrganizer 实例
func NewFolderOrganizer(path string) *FolderOrganizer {
    return &FolderOrganizer{Path: path}
}

// Organize 整理文件夹结构
func (f *FolderOrganizer) Organize(c *gin.Context) {
    // 检查路径是否存在
    if _, err := os.Stat(f.Path); os.IsNotExist(err) {
        c.JSON(400, gin.H{
            "error": "The specified path does not exist.",
        })
        return
    }

    // 获取文件夹中的所有文件和子文件夹
    files, err := os.ReadDir(f.Path)
    if err != nil {
        c.JSON(500, gin.H{
            "error": fmt.Sprintf("Failed to read directory: %s", err),
        })
        return
    }

    // 遍历文件夹中的每个条目
    for _, file := range files {
        fullpath := filepath.Join(f.Path, file.Name())
        // 如果是文件夹，递归整理
        if file.IsDir() {
            f.OrganizeFolder(c, fullpath)
        }
        // 如果是文件，移动到对应的文件夹
        f.MoveFile(c, fullpath)
    }

    c.JSON(200, gin.H{
        "message": "Folder organized successfully.",
    })
}

// OrganizeFolder 递归整理文件夹
func (f *FolderOrganizer) OrganizeFolder(c *gin.Context, path string) {
    // 递归调用 Organize 方法
    f.Organize(c)
}

// MoveFile 移动文件到对应的文件夹
func (f *FolderOrganizer) MoveFile(c *gin.Context, path string) {
    // 这里可以根据文件类型或其他规则来决定文件的新位置
    // 例如，可以根据文件扩展名将文件移动到不同的文件夹
    extension := strings.ToLower(filepath.Ext(path))
    destination := filepath.Join(f.Path, extension[1:] + "s") // 假设根据文件类型创建文件夹

    // 创建目标文件夹，如果它不存在
    if _, err := os.Stat(destination); os.IsNotExist(err) {
        if err := os.Mkdir(destination, 0755); err != nil {
            c.JSON(500, gin.H{
                "error": fmt.Sprintf("Failed to create directory: %s", err),
            })
            return
        }
    }

    // 移动文件到目标文件夹
    if err := os.Rename(path, filepath.Join(destination, filepath.Base(path))); err != nil {
        c.JSON(500, gin.H{
            "error": fmt.Sprintf("Failed to move file: %s", err),
        })
        return
    }
}

func main() {
    r := gin.Default()

    // 创建 FolderOrganizer 实例
    organizer := NewFolderOrganizer("./")

    // 注册路由和处理器
    r.POST("/organize", organizer.Organize)

    // 启动服务器
    r.Run()
}