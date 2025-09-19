// 代码生成时间: 2025-09-19 13:53:57
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
)

// SortAlgorithmService 结构体，用于封装排序算法的逻辑
type SortAlgorithmService struct{}

// BubbleSort 实现冒泡排序算法
func (s *SortAlgorithmService) BubbleSort(arr []int) ([]int, error) {
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr)-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    return arr, nil
}

// SortHandler Gin处理器实现排序逻辑
func SortHandler(c *gin.Context) {
    var service SortAlgorithmService
    var input []int
    if err := c.ShouldBindJSON(&input); err != nil {
        // 错误处理
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }
    
    sorted, err := service.BubbleSort(input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to sort"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"sorted": sorted})
}

func main() {
    r := gin.Default()
    r.POST("/sort", SortHandler) // 绑定POST请求到SortHandler
    r.Run() // 监听并在0.0.0.0:8080上启动服务
}
