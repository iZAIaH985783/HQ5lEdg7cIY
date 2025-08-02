// 代码生成时间: 2025-08-03 00:01:29
package main

import (
    "archive/zip"
    "bytes"
    "io"
    "net/http"
    "os"
    "strings"

    "github.com/gin-gonic/gin" // Gin Web Framework
)

// handleZipUpload is a Gin handler function to handle file upload and unzip
func handleZipUpload(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Error Retrieving the File",
        })
        return
    }

    // Open the uploaded zip file
    src, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Unable to Open the File",
        })
        return
    }
    defer src.Close()

    // Create a buffer to store the contents of the file
    buffer := new(bytes.Buffer)
    _, err = buffer.ReadFrom(src)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Unable to Read the File",
        })
        return
    }

    // Attempt to decompress the zip file
    _, err = zip.NewReader(bytes.NewReader(buffer.Bytes()), int64(len(buffer.Bytes())))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Invalid Zip File",
        })
        return
    }

    // If everything went fine, send a success message
    c.JSON(http.StatusOK, gin.H{
        "message": "File Uploaded and Unzipped Successfully",
    })
}

// main function to run the Gin server
func main() {
    r := gin.Default()

    // Define the route and attach the handler function
    r.POST("/upload", handleZipUpload)

    // Run the server on http://localhost:8080
    r.Run(":8080")
}
