// 代码生成时间: 2025-09-22 23:36:08
package main

import (
    "fmt"
    "net/http"
    "runtime"
    "time"

    "github.com/gin-gonic/gin"
)

// SystemMonitorHandler 定义了一个处理器，用于监控系统性能
type SystemMonitorHandler struct{}

// GetSystemInfo 提供一个HTTP GET请求处理器，返回当前Golang运行时信息
func (h *SystemMonitorHandler) GetSystemInfo(c *gin.Context) {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    info := struct {
        Alloc     uint64 `json:"alloc"`      // 已分配的内存字节数
        TotalAlloc uint64 `json:"total_alloc"` // 总共分配的内存字节数
        Sys       uint64 `json:"sys"`        // 从操作系统获得的内存字节数
        NumGC     uint32 `json:"num_gc"`     // 触发的垃圾回收次数
    }{
        m.Alloc, m.TotalAlloc, m.Sys, m.NumGC,
    }

    c.JSON(http.StatusOK, gin.H{
        "memory_usage": info,
    })
}

// NewSystemMonitorHandler 创建并返回一个新的SystemMonitorHandler实例
func NewSystemMonitorHandler() *SystemMonitorHandler {
    return &SystemMonitorHandler{}
}

func main() {
    r := gin.Default()

    // 实例化系统监控处理器
    monitor := NewSystemMonitorHandler()

    // 注册处理器到路由
    r.GET("/system", monitor.GetSystemInfo)

    // 启动服务
    r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
