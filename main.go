package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "vignesh.com/gocrudrest/config"
	"vignesh.com/gocrudrest/message"
	"vignesh.com/gocrudrest/middleware"
	"vignesh.com/gocrudrest/user"
)

func init() {
	fmt.Println("main init")
}

func main() {
	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(middleware.TrackingIdLogger())

	app.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "running",
		})
	})

	message.Routes(app)
	user.Routes(app)

	app.Run(":8080")
}
