// 代码生成时间: 2025-09-02 18:01:52
package main

import (
    "fmt"
    "math"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

// MathCalculatorHandler 结构体，处理数学计算
type MathCalculatorHandler struct{}

// Add 计算两个数的和
func (handler *MathCalculatorHandler) Add(c *gin.Context) {
    var req struct {
        Num1 float64 `json:"num1" binding:"required"`
        Num2 float64 `json:"num2" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "result": req.Num1 + req.Num2,
    })
}

// Subtract 计算两个数的差
func (handler *MathCalculatorHandler) Subtract(c *gin.Context) {
    var req struct {
        Num1 float64 `json:"num1" binding:"required"`
        Num2 float64 `json:"num2" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "result": req.Num1 - req.Num2,
    })
}

// Multiply 计算两个数的乘积
func (handler *MathCalculatorHandler) Multiply(c *gin.Context) {
    var req struct {
        Num1 float64 `json:"num1" binding:"required"`
        Num2 float64 `json:"num2" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "result": req.Num1 * req.Num2,
    })
}

// Divide 计算两个数的商
func (handler *MathCalculatorHandler) Divide(c *gin.Context) {
    var req struct {
        Num1 float64 `json:"num1" binding:"required"`
        Num2 float64 `json:"num2" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    if req.Num2 == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "division by zero is not allowed",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "result": req.Num1 / req.Num2,
    })
}

// ErrorHandler 中间件函数，用于处理错误
func ErrorHandler(c *gin.Context) {
    c.Next()
    if len(c.Errors) > 0 {
        err := c.Errors.Last().Err
        switch err.(type) {
        case *gin.Error:
            switch v := err.(type) {
            case *gin.Error:
                c.JSON(v.StatusCode, gin.H{
                    "error": v.Err.Error(),
                })
            }
        default:
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "internal server error",
            })
        }
    }
}

func main() {
    router := gin.Default()
    router.Use(ErrorHandler)

    calculator := MathCalculatorHandler{}

    router.POST("/add", calculator.Add)
    router.POST("/subtract", calculator.Subtract)
    router.POST("/multiply", calculator.Multiply)
    router.POST("/divide", calculator.Divide)

    router.Run(":8080")
}
