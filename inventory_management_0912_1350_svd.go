// 代码生成时间: 2025-09-12 13:50:11
package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// InventoryItem 库存项结构体
type InventoryItem struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    Quantity  int    `json:"quantity"`
}

// inventoryItems 模拟库存列表
var inventoryItems = []InventoryItem{
    {ID: "1", Name: "Item1", Quantity: 10},
    {ID: "2", Name: "Item2", Quantity: 20},
}

// InventoryManager 提供库存管理功能
type InventoryManager struct {
}

// AddItem 添加新的库存项
func (im *InventoryManager) AddItem(c *gin.Context) {
    var newItem InventoryItem
    if err := c.ShouldBindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    inventoryItems = append(inventoryItems, newItem)
    c.JSON(http.StatusOK, newItem)
}

// GetItems 获取所有库存项
func (im *InventoryManager) GetItems(c *gin.Context) {
    c.JSON(http.StatusOK, inventoryItems)
}

// UpdateItem 更新指定ID的库存项
func (im *InventoryManager) UpdateItem(c *gin.Context) {
    id := c.Param("id")
    for i, item := range inventoryItems {
        if item.ID == id {
            var updatedItem InventoryItem
            if err := c.ShouldBindJSON(&updatedItem); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
            }
            inventoryItems[i] = updatedItem
            c.JSON(http.StatusOK, updatedItem)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
}

// DeleteItem 删除指定ID的库存项
func (im *InventoryManager) DeleteItem(c *gin.Context) {
    id := c.Param("id")
    for i, item := range inventoryItems {
        if item.ID == id {
            inventoryItems = append(inventoryItems[:i], inventoryItems[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
}

func main() {
    router := gin.Default()

    // 使用中间件记录请求日志
    router.Use(gin.Logger())
    // 使用中间件恢复任何panic，将500错误标准化
    router.Use(gin.Recovery())

    manager := InventoryManager{}

    // 路由组
    v1 := router.Group("/api/v1")
    {"api": "/api/v1"}

        v1.POST("/items", manager.AddItem) // 添加库存项
        v1.GET("/items", manager.GetItems) // 获取所有库存项
        v1.PUT("/items/:id", manager.UpdateItem) // 更新库存项
        v1.DELETE("/items/:id", manager.DeleteItem) // 删除库存项

    fmt.Println("Server starting on :8080")
    router.Run(":8080")
}