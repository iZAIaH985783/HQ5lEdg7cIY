// 代码生成时间: 2025-08-11 01:35:18
package main

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
)

// ConfigManager 结构体用于管理配置文件
type ConfigManager struct {
    viper *viper.Viper
}

// NewConfigManager 创建一个新的ConfigManager实例
func NewConfigManager(configPath string) (*ConfigManager, error) {
    var cm ConfigManager
    cm.viper = viper.New()
    cm.viper.SetConfigType("yaml") // 假设配置文件为YAML格式
    cm.viper.SetConfigName("config") // 配置文件名称（不包含扩展名）
    cm.viper.AddConfigPath(configPath) // 设置配置文件路径

    // 尝试读取配置文件
    if err := cm.viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }

    return &cm, nil
}

// GetConfig 获取配置项的值
func (cm *ConfigManager) GetConfig(key string) interface{} {
    return cm.viper.Get(key)
}

// main 函数，设置路由并启动Gin服务器
func main() {
    // 初始化Gin引擎
    r := gin.Default()

    // 创建配置管理器
    configManager, err := NewConfigManager("./config")
    if err != nil {
        panic(err) // 在实际应用中，应使用更复杂的错误处理逻辑
    }

    // 使用Gin中间件处理静态文件服务，例如服务配置文件所在目录
    r.Static("/config", "./config")

    // 定义一个路由来获取配置信息
    r.GET("/config/:key", func(c *gin.Context) {
        key := c.Param("key")
        configValue := configManager.GetConfig(key)
        if configValue == nil {
            c.JSON(404, gin.H{"error": "Config key not found"})
            return
        }
        c.JSON(200, gin.H{"value": configValue})
    })

    // 启动服务器
    if err := r.Run(":8080"); err != nil {
        fmt.Printf("Failed to start server: %s
", err)
        os.Exit(1)
    }
}
