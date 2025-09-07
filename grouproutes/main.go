package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	// simple group: v1
	{
		v1 := r.Group("/v1")
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	{
		v2 := r.Group("/v2")
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	r.Run(":8080")
}