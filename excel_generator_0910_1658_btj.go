// 代码生成时间: 2025-09-10 16:58:34
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/xuri/excelize/v2"
)

// 初始化Excel文件生成器
func InitExcelGenerator() *excelize.File {
    excelFile := excelize.NewFile()
    // 设置工作表的名称
    sheetName := "Sheet1"
    excelFile.NewSheet(sheetName)
    return excelFile
}

// ExcelGeneratorHandler Gin处理器，用于生成Excel文件
func ExcelGeneratorHandler(c *gin.Context) {
    var data [][]string
    // 示例数据，实际应用中应从数据库或其他来源获取
    data = append(data, []string{"Name", "Age"})
    data = append(data, []string{"John", "30"})
    data = append(data, []string{"Jane", "25"})

    // 创建Excel文件
    excelFile := InitExcelGenerator()
    // 设置单元格的值
    for i, row := range data {
        for j, val := range row {
            err := excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+1), val)
            if err != nil {
                // 错误处理
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error": "Failed to set cell value",
                })
                return
            }
        }
    }
    // 保存Excel文件
    filePath := "example.xlsx"
    if err := excelFile.SaveAs(filePath); err != nil {
        // 错误处理
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to save Excel file",
        })
        return
    }

    // 设置响应头
    c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
    c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filePath))
    c.File(filePath)
}

func main() {
    r := gin.Default()
    // 可以添加其他中间件
    r.Use(gin.Recovery())
    r.GET("/excel", ExcelGeneratorHandler)
    // 启动服务
    log.Printf("Server is running on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Error starting server: %v", err)
    }
}
