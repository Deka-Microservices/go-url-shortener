package routes

import (
	"math/rand"
	"net/http"

	"github.com/deka-microservices/go-url-shortener/internal/consts"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Health(ctx *gin.Context) {
	rng := rand.Uint32() % 100
	if rng < viper.GetUint32(consts.CONFIG_FAIL_PERCENT) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fake bad helath",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
