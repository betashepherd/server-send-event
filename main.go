package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/stream", func(c *gin.Context) {
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")
		c.Header("Transfer-Encoding", "chunked")
		c.Header("Access-Control-Allow-Origin", "*")

		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ticker.C:
				data := "Hello " + time.Now().Format("2006-01-02 15:04:05") + "\n\n"
				fmt.Fprintf(c.Writer, "data: %s\n\n", data)
				c.Writer.Flush()
			case <-c.Writer.CloseNotify():
				// If the client terminates the connection, stop the ticker
				ticker.Stop()
				return
			}
		}
	})

	// Parse Static files
	r.StaticFile("/", "./public/index.html")

	// 启动服务器
	r.Run(":8080")
}
