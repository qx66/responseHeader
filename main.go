package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("11")
	r := gin.Default()
	r.GET("/*any", reply)
	panic(r.Run(":8081"))
}

func reply(c *gin.Context) {
	headers := c.Request.Header
	for k, v := range headers {
		if len(v) == 0 {
			continue
		}
		c.Header(k, v[0])
	}
	
	headerSize := 0
	for k, vals := range c.Writer.Header() {
		for _, val := range vals {
			headerSize += len(k) + 2 + len(val) + 2 // ": " + "\r\n"
		}
	}
	c.JSON(200, gin.H{
		"headerSize": headerSize,
		"headerUnit": "Byte",
	})
}
