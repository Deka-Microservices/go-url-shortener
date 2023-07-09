package routes

import (
	"errors"
	"net/http"

	"github.com/deka-microservices/go-url-shortener/internal/database"
	"github.com/deka-microservices/go-url-shortener/internal/database/dberrors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GetLong(ctx *gin.Context) {
	url := ctx.Param("url")

	longUrl, err := database.DB.Get(url)
	if err != nil {
		if errors.Is(err, dberrors.ErrKeyHasAlreadyExists) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "unknown url",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		}
	}

	log.Info().Str("short_url", url).Str("long_url", longUrl).Msg("resolved short url")

	ctx.Redirect(http.StatusMovedPermanently, longUrl)
}
