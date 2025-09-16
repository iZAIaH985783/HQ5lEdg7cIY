// 代码生成时间: 2025-09-17 01:50:01
package main

import (
    "fmt"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
)

// ResponseData is a generic structure to return consistent API responses
type ResponseData struct {
    Data      interface{} `json:"data"`
    Message   string     `json:"message"`
    Error     string     `json:"error"`
    Timestamp int64      `json:"timestamp"`
}

// ErrorResponse is a structure to return error responses
type ErrorResponse struct {
    ResponseData
}

// SuccessResponse is a structure to return success responses
type SuccessResponse struct {
    ResponseData
}

// NewErrorResponse creates a new ErrorResponse
func NewErrorResponse(err error) ErrorResponse {
    return ErrorResponse{
        ResponseData: ResponseData{
            Error:     err.Error(),
            Timestamp: time.Now().Unix(),
        },
    }
}

// NewSuccessResponse creates a new SuccessResponse
func NewSuccessResponse(data interface{}, message string) SuccessResponse {
    return SuccessResponse{
        ResponseData: ResponseData{
            Data:      data,
            Message:   message,
            Timestamp: time.Now().Unix(),
        },
    }
}

// APIResponseFormatter is a middleware that formats API responses
func APIResponseFormatter(c *gin.Context) {
    c.Next()
    response := c.Get(gin.Keys{key: "response"})
    if response == nil {
        // In case no response was set, we set a generic success response
        c.JSON(http.StatusOK, NewSuccessResponse(nil, "No data provided"))
    } else {
        c.JSON(http.StatusOK, response)
    }
}

// NewRouter initializes the router with middleware and routes
func NewRouter() *gin.Engine {
    router := gin.Default()
    router.Use(gin.Recovery())
    router.Use(APIResponseFormatter)
    // Add routes here using router.HandleFunc
    return router
}

// main function to start the HTTP server
func main() {
    router := NewRouter()
    // Define routes
    router.GET("/example", func(c *gin.Context) {
        // Set a custom response
        c.Set(gin.Keys{key: "response"}, NewSuccessResponse("Hello, World!", "Example Success Response"))
    })
    // Start the server
    router.Run()
}
