// 代码生成时间: 2025-08-12 07:33:19
package main

import (
    "archive/zip"
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

// FileDecompressService 结构体用于封装解压文件所需的方法
type FileDecompressService struct {
    OutputPath string
}

// NewFileDecompressService 创建一个新的FileDecompressService实例
func NewFileDecompressService(outputPath string) *FileDecompressService {
    return &FileDecompressService{
        OutputPath: outputPath,
    }
}

// Decompress 用于解压文件到指定路径
func (s *FileDecompressService) Decompress(r *http.Request) error {
    // 从请求中读取文件
    file, header, err := r.FormFile("file")
    if err != nil {
        return err
    }
    defer file.Close()

    // 创建缓冲区以存储文件数据
    var buf bytes.Buffer
    _, err = io.Copy(&buf, file)
    if err != nil {
        return err
    }
    src := buf.Bytes()

    // 创建一个zip reader
    zr, err := zip.NewReader(bytes.NewReader(src), int64(len(src)))
    if err != nil {
        return err
    }
    defer zr.Close()

    // 遍历zip文件中的每个文件
    for _, f := range zr.File {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()

        // 确保目标路径存在
        if err := os.MkdirAll(s.OutputPath, 0755); err != nil {
            return err
        }

        // 构建目标文件路径
        targetPath := filepath.Join(s.OutputPath, f.Name)
        fmt.Printf("Writing to: %s
", targetPath)

        // 创建目标文件
        w, err := os.Create(targetPath)
        if err != nil {
            return err
        }
        defer w.Close()

        // 将zip文件内容写入目标文件
        _, err = io.Copy(w, rc)
        if err != nil {
            return err
        }
    }
    return nil
}

func main() {
    router := gin.Default()

    // 创建解压服务实例
    decompressService := NewFileDecompressService("./decompressed")

    // 定义一个路由来处理文件上传和解压
    router.POST("/decompress", func(c *gin.Context) {
        // 调用解压服务方法
        if err := decompressService.Decompress(c.Request); err != nil {
            // 如果发生错误，返回错误信息
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "message": "File decompressed successfully",
        })
    })

    // 启动服务器
    router.Run(":8080")
}
