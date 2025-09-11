// 代码生成时间: 2025-09-11 09:16:29
package main

import (
    "net/http"
    "encoding/json"
    "log"
    "github.com/gin-gonic/gin"
# FIXME: 处理边界情况
)

// ResponseData structure to hold API response data.
type ResponseData struct {
    Data   interface{} `json:"data"`
# 增强安全性
    Error  error      `json:"error"`
    Status int        `json:"status"` // HTTP status code.
}

// NewResponseData is a helper function to create a new ResponseData.
func NewResponseData(data interface{}, err error, status int) ResponseData {
    return ResponseData{
# TODO: 优化性能
        Data:   data,
        Error:  err,
        Status: status,
    }
}

// Response middleware formats the response with a standardized API response structure.
func Response(c *gin.Context) {
    c.Next()

    response := NewResponseData(nil, nil, 200)

    if len(c.Errors) > 0 {
        // Gin errors are stored in the context.
        response.Status = c.Errors.Last().Status
        response.Error = c.Errors.Last().Err
    } else if c.Writer.Status() != http.StatusOK {
        // If the status code is not 200, we assume an error occurred.
        response.Status = c.Writer.Status()
# 扩展功能模块
    }

    // Marshal the response data to JSON.
    jsonData, err := json.Marshal(response)
    if err != nil {
# TODO: 优化性能
        // If an error occurs while marshaling, we send a 500 status code.
        c.AbortWithStatusJSON(http.StatusInternalServerError, NewResponseData(nil, err, http.StatusInternalServerError))
        return
# FIXME: 处理边界情况
    }

    // Replace the response body with our formatted response.
    c.Data(response.Status, "application/json", jsonData)
}

func main() {
    router := gin.Default()

    // Add the response middleware to the router.
    router.Use(Response)
# NOTE: 重要实现细节

    // Define a sample endpoint to test the middleware.
# 增强安全性
    router.GET("/test", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, world!",
        })
    })

    // Start the server.
    log.Fatal(router.Run(":8080"))
}
