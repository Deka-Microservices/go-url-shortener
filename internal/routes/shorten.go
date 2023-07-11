package routes

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"

	"github.com/deka-microservices/go-url-shortener/internal/database"
	"github.com/deka-microservices/go-url-shortener/pkg/base62"
	"github.com/deka-microservices/go-url-shortener/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Shorten(ctx *gin.Context) {
	var req models.ShortenRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	url_regexp := "^https?:\\/\\/(?:www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b(?:[-a-zA-Z0-9()@:%_\\+.~#?&\\/=]*)$"
	matched, err := regexp.Match(url_regexp, []byte(req.Url))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if !matched {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "url format is invalid",
		})
		return
	}

	n := rand.Uint32()
	s := base62.Encode(n)

	exists, _ := database.DB.Exists(ctx.Request.Context(), s)
	fail_count := 16
	for exists {
		fail_count--
		if fail_count <= 0 {
			ctx.JSON(http.StatusInsufficientStorage, gin.H{
				"message": "there are no place to store new short url",
			})
			return
		}

		n = rand.Uint32()
		s = base62.Encode(n)
		exists, _ = database.DB.Exists(ctx.Request.Context(), s)
	}

	log.Info().Str("short_url", s).Str("long_url", req.Url).Msg("short_url_generated")
	database.DB.AddUrl(ctx.Request.Context(), s, req.Url)

	ctx.JSON(http.StatusCreated, gin.H{
		"short_url": fmt.Sprintf("/%s", s),
	})
}
