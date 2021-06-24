package message

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Debug("msg controller init")
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
