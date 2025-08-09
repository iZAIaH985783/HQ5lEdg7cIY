// 代码生成时间: 2025-08-09 11:58:58
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
)

// FormValidator 结构体包含表单验证所需的字段
type FormValidator struct {
    Username string `form:"username" binding:"required,min=3,max=15" json:"username"`
    Email    string `form:"email" binding:"required,email" json:"email"`
    Age      int    `form:"age" binding:"required,gte=18,lte=99" json:"age"`
}

// ValidateForm 处理表单数据并进行验证
func ValidateForm(c *gin.Context) {
    var form FormValidator
    // 使用 ShouldBindForm 进行表单数据绑定和验证
    if err := c.ShouldBindForm(&form); err != nil {
        // 如果验证失败，返回错误信息
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    // 表单验证通过，返回成功信息
    c.JSON(http.StatusOK, gin.H{
        "message": "Form data is valid",
        "data": form,
    })
}

func main() {
    // 创建一个新的Gin路由器
    router := gin.Default()

    // 注册表单验证处理器
    router.POST("/form", ValidateForm)

    // 启动服务器
    router.Run(":8080")
}
