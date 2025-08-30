// 代码生成时间: 2025-08-31 06:30:01
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// MathCalculatorHandler 结构体，用于处理数学计算请求
type MathCalculatorHandler struct {
}

// AddHandler 处理加法请求
func (h *MathCalculatorHandler) AddHandler(c *gin.Context) {
    a, b := float64(c.PostForm("a")), float64(c.PostForm("b"))
    if a < 0 || b < 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Negative numbers are not allowed",
        })
        return
    }
    sum := a + b
    c.JSON(http.StatusOK, gin.H{
        "result": sum,
    })
}

// SubtractHandler 处理减法请求
func (h *MathCalculatorHandler) SubtractHandler(c *gin.Context) {
    a, b := float64(c.PostForm("a")), float64(c.PostForm("b"))
    if a < 0 || b < 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Negative numbers are not allowed",
        })
        return
    }
    difference := a - b
    c.JSON(http.StatusOK, gin.H{
        "result": difference,
    })
}

// MultiplyHandler 处理乘法请求
func (h *MathCalculatorHandler) MultiplyHandler(c *gin.Context) {
    a, b := float64(c.PostForm("a")), float64(c.PostForm("b"))
    if a < 0 || b < 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Negative numbers are not allowed",
        })
        return
    }
    product := a * b
    c.JSON(http.StatusOK, gin.H{
        "result": product,
    })
}

// DivideHandler 处理除法请求
func (h *MathCalculatorHandler) DivideHandler(c *gin.Context) {
    a, b := float64(c.PostForm("a")), float64(c.PostForm("b"))
    if a < 0 || b < 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Negative numbers are not allowed",
        })
        return
    }
    if b == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Division by zero is not allowed",
        })
        return
    }
    quotient := a / b
    c.JSON(http.StatusOK, gin.H{
        "result": quotient,
    })
}

func main() {
    r := gin.Default()

    // 创建一个MathCalculatorHandler实例
    calculator := MathCalculatorHandler{}

    // 路由到加法处理函数
    r.POST("/add", calculator.AddHandler)

    // 路由到减法处理函数
    r.POST("/subtract", calculator.SubtractHandler)

    // 路由到乘法处理函数
    r.POST="/multiply", calculator.MultiplyHandler)

    // 路由到除法处理函数
    r.POST("/divide", calculator.DivideHandler)

    // 启动服务器
    r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
