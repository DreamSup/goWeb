package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world!")
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!",
		})
	})
	app.Run(":8080")

}
