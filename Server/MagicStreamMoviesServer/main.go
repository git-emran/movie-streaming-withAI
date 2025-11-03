package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "wooow, Its serving time")
	})

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start the serving")

	}

}
