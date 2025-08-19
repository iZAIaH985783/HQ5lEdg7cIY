// 代码生成时间: 2025-08-19 20:55:38
package main

import (
    "fmt"
    "net/http"
    "os/exec"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/gin-gonic/gin"
)

// ProcessManager 定义进程管理器结构
type ProcessManager struct {
    Command *exec.Cmd
}

// NewProcessManager 创建新的进程管理器
func NewProcessManager(cmd *exec.Cmd) *ProcessManager {
    return &ProcessManager{Command: cmd}
}

// Start 启动进程
func (pm *ProcessManager) Start() error {
    return pm.Command.Start()
}

// Stop 停止进程
func (pm *ProcessManager) Stop() error {
    if pm.Command.Process == nil {
        return fmt.Errorf("process is not running")
    }
    return pm.Command.Process.Signal(syscall.SIGTERM)
}

// GinProcessManagerRoute 设置Gin路由和中间件来管理进程
func GinProcessManagerRoute(router *gin.Engine) {
    pm := NewProcessManager(exec.Command("your_command_here"))
    // 启动进程
    if err := pm.Start(); err != nil {
        fmt.Println("Failed to start process: ", err)
        os.Exit(1)
    }
    defer pm.Stop()

    // 设置Gin中间件
    router.Use(gin.Recovery())

    // 进程状态端点
    router.GET("/process/status", func(c *gin.Context) {
        if pm.Command.Process != nil && pm.Command.ProcessState != nil && pm.Command.ProcessState.Running() {
            c.JSON(http.StatusOK, gin.H{
                "status": "running",
            })
        } else {
            c.JSON(http.StatusOK, gin.H{
                "status": "stopped",
            })
        }
    })

    // 启动和停止进程的端点
    router.POST("/process/start", func(c *gin.Context) {
        if err := pm.Start(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "status": "started",
        })
    })

    router.POST("/process/stop", func(c *gin.Context) {
        if err := pm.Stop(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "status": "stopped",
        })
    })
}

// main 函数设置Gin路由器并启动服务器
func main() {
    router := gin.Default()
    GinProcessManagerRoute(router)

    // 捕获中断信号以优雅地关闭进程和服务器
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigs
        // 这里可以添加更多的清理代码
        fmt.Println("Shutting down...")
        // 停止Gin服务器
        router.ShutdownNow()
    }()

    // 启动Gin服务器
    fmt.Println("Server started at :8080")
    if err := router.Run(":8080"); err != nil {
        fmt.Println("Failed to start server: ", err)
    }
}