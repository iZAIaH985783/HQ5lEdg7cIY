// 代码生成时间: 2025-09-23 20:18:32
package main

import (
    "fmt"
    "net/http"
    "runtime"
    "time"

    "github.com/gin-gonic/gin"
)

// MemoryUsageHandler 处理器用于分析内存使用情况
func MemoryUsageHandler(c *gin.Context) {
    // 获取当前内存使用情况
    var m runtime.MemStats
    runtime.ReadMemStats(&m)

    // 内存使用情况响应对象
    memoryUsage := struct {
        Alloc       uint64    `json:"alloc"`       // 已分配的堆内存
        TotalAlloc uint64    `json:"totalAlloc"` // 到当前时刻为止分配的总内存
        Sys         uint64    `json:"sys"`         // 从操作系统获得的内存
        NumGC       uint32    `json:"numGC"`       // 触发的垃圾回收次数
    }{
        Alloc: m.Alloc,
        TotalAlloc: m.TotalAlloc,
        Sys: m.Sys,
        NumGC: m.NumGC,
    }

    // 将内存使用情况写入响应
    c.JSON(http.StatusOK, memoryUsage)
}

func main() {
    r := gin.Default()

    // 注册内存使用情况处理器
    r.GET("/memory", MemoryUsageHandler)

    // 启动服务器
    r.Run()
}
