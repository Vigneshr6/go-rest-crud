package main

import (
	"github.com/gin-gonic/gin"
	"vignesh.com/gocrudrest/message"
	"vignesh.com/gocrudrest/user"
)

func main() {
	app := gin.Default()

	app.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "running",
		})
	})

	message.Routes(app)
	user.Routes(app)

	app.Run(":8080")
}
