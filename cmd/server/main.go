package main

import (
	_ "github.com/deka-microservices/go-url-shortener/internal/config"
	"github.com/deka-microservices/go-url-shortener/internal/consts"
	"github.com/deka-microservices/go-url-shortener/internal/database"
	"github.com/deka-microservices/go-url-shortener/internal/routes"
	"github.com/gin-contrib/logger"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(logger.SetLogger())

	// DB
	database.InitGlobalDatabase()

	// Init routes
	e.GET("/:url", routes.GetLong)
	e.POST("/shorten", routes.Shorten)

	address := "0.0.0.0:" + viper.GetString(consts.CONFIG_PORT)

	log.Info().Str("version", consts.Version()).Msg("version report")

	e.Run(address)
}
