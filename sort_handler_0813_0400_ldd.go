// 代码生成时间: 2025-08-13 04:00:20
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "sort"
)

// SortHandler 结构体包含排序算法的实现
type SortHandler struct {
    // 可以添加更多属性，例如排序算法的选择
}

// NewSortHandler 返回一个新的SortHandler实例
func NewSortHandler() *SortHandler {
    return &SortHandler{}
}

// SortNumbers 处理数字列表的排序
// 该方法接受一个整数切片并返回一个排序后的切片
func (s *SortHandler) SortNumbers(numbers []int) ([]int, error) {
    if numbers == nil {
        return nil, ErrorNilInput
