// 代码生成时间: 2025-09-10 12:06:03
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gin-gonic/gin"
)

// ErrorLogger is a struct that contains information about the request and response.
type ErrorLogger struct {
    logFileName string
}

// NewErrorLogger creates a new ErrorLogger instance with a log file name.
func NewErrorLogger(logFileName string) *ErrorLogger {
    return &ErrorLogger{
        logFileName: logFileName,
    }
}

// ErrorHandling is the middleware function that logs errors and writes the response.
func (l *ErrorLogger) ErrorHandling(c *gin.Context) {
    recover := recover()
    if recover != nil {
        log.Printf("Recovered in %v", c.Request.URL)
        // Log error details to the file
        errorLog := fmt.Sprintf("%v - [%v] %v\
", time.Now().Format(time.RFC3339), c.Request.URL, recover)
        f, err := os.OpenFile(l.logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
        if err != nil {
            // Handle file opening error
            log.Printf("Error opening log file: %v", err)
        } else {
            if _, err := f.WriteString(errorLog); err != nil {
                // Handle error writing to the file
                log.Printf("Error writing to log file: %v", err)
            }
            f.Close()
        }
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "Internal Server Error", "error": fmt.Sprintf("%+v", recover),
        })
    }
}

func main() {
    r := gin.Default()
    logPath := "./error.log"
    errorLogger := NewErrorLogger(logPath)
    r.Use(errorLogger.ErrorHandling)

    // Example route that will trigger an error
    r.GET("/error", func(c *gin.Context) {
        panic("Something went wrong!")
    })

    // Start the server
    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}