// 代码生成时间: 2025-09-22 14:59:42
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/robfig/cron/v3"
)

// SchedulerWithGin demonstrates a simple Gin-based HTTP server with a built-in cron scheduler.
func main() {
    // Create a new Gin router with default middleware.
    r := gin.Default()

    // Set up a cron scheduler.
    schedule, err := cron.New(cron.WithSeconds())
    if err != nil {
        log.Fatalf("Error setting up cron scheduler: %v", err)
    }
    defer schedule.Stop()

    // Define a job to run a task every minute.
    schedule.AddFunc("@every 1m", func() { handleScheduledTask() })

    // Start the scheduler.
    go schedule.Start()

    // Define an endpoint to trigger a manual task run.
    r.GET("/run-task", func(c *gin.Context) {
        if err := runTask(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
        } else {
            c.JSON(http.StatusOK, gin.H{
                "message": "Task ran successfully.",
            })
        }
    })

    // Start the HTTP server.
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}

// runTask is a function that simulates a task.
func runTask() error {
    // Simulate some work by sleeping for a second.
    time.Sleep(1 * time.Second)
    return nil
}

// handleScheduledTask is a function that runs when the cron job is triggered.
func handleScheduledTask() {
    fmt.Println("Scheduled task is running...")
    if err := runTask(); err != nil {
        fmt.Printf("Error running scheduled task: %v", err)
    } else {
        fmt.Println("Scheduled task completed successfully.")
    }
}
