// 代码生成时间: 2025-08-12 20:29:48
package main

import (
    "fmt"
    "log"
    "net/http"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
)

// DBPool 定义数据库连接池
type DBPool struct {
    *gorm.DB
}

func main() {
    r := gin.Default()

    // 初始化数据库连接池
    dbPool, err := NewDBPool()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer dbPool.Close()

    // 定义中间件，用于获取数据库连接
    dbPoolMiddleware := func(c *gin.Context) {
        c.Set("db", dbPool)
        c.Next()
    }

    // 注册中间件
    r.Use(dbPoolMiddleware)

    // 定义路由和处理器
    r.GET("/", func(c *gin.Context) {
        db := c.MustGet("db").(*DBPool)
        // 这里可以添加使用数据库的逻辑
        fmt.Println("Database connection acquired")
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    // 启动服务
    r.Run()
}

// NewDBPool 初始化并返回一个新的数据库连接池
func NewDBPool() (*DBPool, error) {
    dsn := "file:mydb.sqlite?cache=shared&mode=rwc"
    db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    // 自动迁移以确保数据库结构是最新的
    db.AutoMigrate()
    return &DBPool{DB: db}, nil
}
