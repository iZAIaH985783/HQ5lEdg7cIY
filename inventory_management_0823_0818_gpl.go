// 代码生成时间: 2025-08-23 08:18:24
package main

import (
    "fmt"
    "net/http"
    "log"

    "github.com/gin-gonic/gin"
)

// InventoryItem 库存项
type InventoryItem struct {
    ID          string `json:"id"`
    Description string `json:"description"`
    Quantity    int    `json:"quantity"`
}

// InventoryHandler 库存处理器
type InventoryHandler struct {
    items map[string]InventoryItem
}

// NewInventoryHandler 创建一个新的InventoryHandler实例
func NewInventoryHandler() *InventoryHandler {
    return &InventoryHandler{
        items: make(map[string]InventoryItem),
    }
}

// AddItem 添加库存项
func (h *InventoryHandler) AddItem(c *gin.Context) {
    var item InventoryItem
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": fmt.Sprintf("Invalid input: %v", err),
        })
        return
    }
    h.items[item.ID] = item
    c.JSON(http.StatusOK, gin.H{
        "status":    "success",
        "item":      item,
        "message":   "Item added successfully",
    })
}

// GetItem 获取库存项
func (h *InventoryHandler) GetItem(c *gin.Context) {
    id := c.Param("id\)
    item, exists := h.items[id]
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{
            "error": fmt.Sprintf("Item with ID %s not found", id),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "item":    item,
    })
}

// UpdateItem 更新库存项
func (h *InventoryHandler) UpdateItem(c *gin.Context) {
    id := c.Param("id")
    var item InventoryItem
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": fmt.Sprintf("Invalid input: %v", err),
        })
        return
    }
    if _, exists := h.items[id]; !exists { {
        return
    }
    h.items[id] = item
    c.JSON(http.StatusOK, gin.H{
        "status":    "success",
        "item":      item,
        "message":   "Item updated successfully",
    })
}

// DeleteItem 删除库存项
func (h *InventoryHandler) DeleteItem(c *gin.Context) {
    id := c.Param("id")
    if _, exists := h.items[id]; !exists {
        c.JSON(http.StatusNotFound, gin.H{
            "error": fmt.Sprintf("Item with ID %s not found", id),
        })
        return
    }
    delete(h.items, id)
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Item deleted successfully",
    })
}

func main() {
    r := gin.Default()

    // 创建库存处理器
    inventory := NewInventoryHandler()

    // 路由设置
    r.POST("/items", inventory.AddItem)
    r.GET("/items/:id", inventory.GetItem)
    r.PUT("/items/:id", inventory.UpdateItem)
    r.DELETE("/items/:id", inventory.DeleteItem)

    // 启动服务器
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
