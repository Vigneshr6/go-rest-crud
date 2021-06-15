package message

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("msg controller init")
}

func Routes(app *gin.Engine) {
	route := app.Group("/msg")
	route.GET("/", getMessage)
}

func getMessage(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello",
	})
}
