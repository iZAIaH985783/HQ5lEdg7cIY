// 代码生成时间: 2025-08-03 17:51:27
package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

func themeSwitchHandler(c *gin.Context) {
    // 获取请求参数
    theme := c.DefaultQuery("theme", "default")
    
    // 错误处理
    if theme != "dark" && theme != "light" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid theme. Please choose 'dark' or 'light'.",
        })
        return
    }
    
    // 保存主题选择到客户端，这里作为示例使用cookie
    c.SetCookie("theme", theme, 3600, "/", "", false, true)
    
    // 响应客户端
    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("Theme switched to %s", theme),
    })
}

func main() {
    // 创建一个新的Gin路由器
    router := gin.Default()
    
    // 路由设置
    router.GET("/theme", themeSwitchHandler)
    
    // 启动服务器
    router.Run(":8080")
}
