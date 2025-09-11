// 代码生成时间: 2025-09-12 01:14:35
// image_resizer.go
// This file contains a Gin-Gonic handler that resizes images in bulk.

package main

import (
    "encoding/json"
    "image"
    "image/jpeg"
    "net/http"
    "os"
    "path/filepath"
    "github.com/gin-gonic/gin"
    "github.com/disintegration/imaging"
)

// ErrorResponse defines the structure of the response sent in case of an error.
type ErrorResponse struct {
    Error string `json:"error"`
}

// ImageResizer holds the configuration for resizing images.
type ImageResizer struct {
    Width  int
    Height int
}

// NewImageResizer returns a new ImageResizer with the given width and height.
func NewImageResizer(width, height int) *ImageResizer {
    return &ImageResizer{
        Width:  width,
        Height: height,
    }
}

// ResizeImage resizes an image to the specified dimensions.
func (r *ImageResizer) ResizeImage(imagePath string) (string, error) {
    img, err := imaging.Open(imagePath)
    if err != nil {
        return "", err
    }
    
    resizedImg := imaging.Resize(img, r.Width, r.Height, imaging.Lanczos)
    
    resizedPath := imagePath + "_resized"
    if err := jpeg.Encode(&os.File{}, resizedImg, nil); err != nil {
        return "", err
    }
    
    return resizedPath, nil
}

// HandleResize is a Gin handler that resizes images in a directory.
func HandleResize(resizer *ImageResizer) gin.HandlerFunc {
    return func(c *gin.Context) {
        dirPath := c.Query("dir")
        if dirPath == "" {
            c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Directory path is required"})
            return
        }
        
        dir, err := os.Open(dirPath)
        if err != nil {
            c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to open directory"})
            return
        }
        defer dir.Close()
        
        files, err := dir.Readdir(-1)
        if err != nil {
            c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to read directory contents"})
            return
        }
        
        resizedPaths := []string{}
        for _, file := range files {
            if !file.IsDir() && (filepath.Ext(file.Name()) == ".jpg" || filepath.Ext(file.Name()) == ".jpeg") {
                imagePath := filepath.Join(dirPath, file.Name())
                resizedPath, err := resizer.ResizeImage(imagePath)
                if err != nil {
                    c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to resize image"})
                    return
                }
                resizedPaths = append(resizedPaths, resizedPath)
            }
        }
        
        c.JSON(http.StatusOK, gin.H{
            "resized_images": resizedPaths,
        })
    }
}

func main() {
    router := gin.Default()
    
    // Middleware to handle logging
    router.Use(gin.Logger())
    
    // Middleware to handle recovery from panics
    router.Use(gin.Recovery())
    
    resizer := NewImageResizer(800, 600)
    router.GET("/resize", HandleResize(resizer))
    
    router.Run(":8080") // listen and serve on 0.0.0.0:8080
}