// 代码生成时间: 2025-08-20 21:02:03
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

// 定义数据库模型
type User struct {
    gorm.Model
    Name      string
    Email     string `gorm:"type:varchar(100);uniqueIndex"`
    Age      int
    Password string
}

// 配置Gin中间件
func setupRouter() *gin.Engine {
    router := gin.Default()

    // 限制请求体大小
    router.MaxMultipartMemory = 8 << 20 // 8 MB

    return router
}

// 使用预编译的SQL语句防止SQL注入
func getUser(c *gin.Context) {
    var user User
    var err error

    // 从请求中获取参数
    id := c.Param("id")

    // 使用预编译的SQL语句和参数化查询
    // 这样可以防止SQL注入
    db, err := getDBConnection()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Database connection failed",
        })
        return
    }
    result := db.First(&user, id)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": result.Error.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, user)
}

// 错误处理中间件
func errorHandlingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next() // 处理请求
        if len(c.Errors) > 0 {
            // 如果有错误，处理它们
            for _, e := range c.Errors {
                c.JSON(http.StatusBadRequest, gin.H{
                    "error": e.Err.Error(),
                })
            }
        }
    }
}

// getDBConnection 用于获取数据库连接
func getDBConnection() (*gorm.DB, error) {
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}

func main() {
    router := setupRouter()
    router.Use(errorHandlingMiddleware()) // 使用错误处理中间件
    router.GET("/user/:id", getUser)
    router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
