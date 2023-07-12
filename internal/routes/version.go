package routes

import (
	"net/http"

	"github.com/deka-microservices/go-url-shortener/internal/consts"
	"github.com/gin-gonic/gin"
)

func Version(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"version": consts.Version(),
	})
}
