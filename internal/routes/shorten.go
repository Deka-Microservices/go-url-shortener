package routes

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/deka-microservices/go-url-shortener/internal/database"
	"github.com/deka-microservices/go-url-shortener/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Shorten(ctx *gin.Context) {
	var req models.ShortenRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
	}

	// TODO! validate link

	n := rand.Uint64()
	s := strconv.FormatUint(n, 16)

	exists, _ := database.DB.Exists(s)
	fail_count := 16
	for exists {
		fail_count--
		if fail_count <= 0 {
			ctx.JSON(http.StatusInsufficientStorage, gin.H{
				"message": "there are no place to store new short url",
			})
			return
		}

		n = rand.Uint64()
		s = strconv.FormatUint(n, 16)
		exists, _ = database.DB.Exists(s)
	}

	log.Info().Str("short_url", s).Str("long_url", req.Url).Msg("short_url_generated")
	database.DB.AddUrl(s, req.Url)

	ctx.JSON(http.StatusCreated, gin.H{
		"short_url": fmt.Sprintf("/%s", s),
	})
}
