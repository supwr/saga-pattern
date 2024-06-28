package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	api := gin.Default()

	api.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Payment OK")
	})

	api.Run(":8002")
}
