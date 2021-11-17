package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"vignesh.com/gocrudrest/common/log"
)

func TrackingIdLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.SetMdc(ctx, "trackingId", uuid.New())
		ctx.Next()
	}
}
