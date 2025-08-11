// 代码生成时间: 2025-08-11 20:40:50
package main

import (
    "archive/zip"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "github.com/gin-gonic/gin"
)

// unzipHandler 是处理解压文件的函数
func unzipHandler(c *gin.Context) {
    // 获取上传的文件
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "No file part"
        })
        return
    }

    // 打开文件
    src, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Could not open uploaded file"
        })
        return
    }
    defer src.Close()

    // 创建一个读取器
    r, err := zip.OpenReader(file.Filename)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Invalid zip file"
        })
        return
    }
    defer r.Close()

    // 创建目标目录
    destPath := filepath.Join(".", "unzipped")
    if _, err := os.Stat(destPath); os.IsNotExist(err) {
        if err := os.MkdirAll(destPath, 0755); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Could not create destination directory"
            })
            return
        }
    }

    // 解压文件
    for _, f := range r.File {
        rc, err := f.Open()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Could not open file inside zip"
            })
            return
        }
        defer rc.Close()

        // 创建目标文件路径
        destFilePath := filepath.Join(destPath, f.Name)
        if f.FileInfo().IsDir() {
            // 如果是目录，则创建目录
            if err := os.MkdirAll(destFilePath, f.Mode()); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error": "Could not create directory"
                })
                return
            }
        } else {
            // 如果是文件，则创建文件
            file, err := os.OpenFile(destFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, f.Mode())
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error": "Could not create file"
                })
                return
            }
            defer file.Close()

            // 复制文件内容
            _, err = io.Copy(file, rc)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error": "Could not copy file content"
                })
                return
            }
        }
    }

    // 返回解压成功消息
    c.JSON(http.StatusOK, gin.H{
        "message": "File successfully unzipped"
    })
}

func main() {
    r := gin.Default()
    r.POST("/unzip", unzipHandler)
    r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
