// 代码生成时间: 2025-08-25 17:06:42
package main

import (
    "net/http"
    "strings"
    "golang.org/x/net/html"
    "github.com/gin-gonic/gin"
)

// struct to hold the scraping result
type ScrapeResult struct {
    Content string `json:"content"`
}

func main() {
    r := gin.Default()

    // Register a route for scraping web content
    r.GET("/scrape", func(c *gin.Context) {
        url := c.Query("url")
        if url == "" {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "URL parameter is required",
            })
            return
        }

        // Make the HTTP GET request to the provided URL
        resp, err := http.Get(url)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to fetch URL",
            })
            return
        }
        defer resp.Body.Close()
        
        // Parse the HTML content
        node, err := html.Parse(resp.Body)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to parse HTML",
            })
            return
        }

        // Extract text content from the HTML
        var scraper func(*html.Node)
        scraper = func(n *html.Node) {
            switch n.Type {
            case html.TextNode:
                text := strings.TrimSpace(n.Data)
                if len(text) > 0 {
                    // Append the text to the result
                    c.Set("result", (c.GetString("result") + " " + text))
                }
            case html.ElementNode:
                for child := n.FirstChild; child != nil; child = child.NextSibling {
                    scraper(child)
                }
            }
        }
        scraper(node)

        // Get the result and send it as JSON response
        result := c.GetString("result")
        c.JSON(http.StatusOK, ScrapeResult{Content: result})
    })

    // Start the server
    r.Run(":8080") // listening and serving on 0.0.0.0:8080
}
