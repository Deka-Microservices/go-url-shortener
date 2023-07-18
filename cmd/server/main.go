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
	e.GET("/health", routes.Health)
	e.GET("/version", routes.Version)
	e.GET("/:url", routes.GetLong)
	e.POST("/shorten", routes.Shorten)

	address := "0.0.0.0:" + viper.GetString(consts.CONFIG_PORT)

	log.Info().Str("version", consts.Version()).Msg("version report")

	certFile := viper.GetString(consts.CONFIG_TLS_CERT)
	keyFile := viper.GetString(consts.CONFIG_TLS_KEY)

	if len(certFile) != 0 && len(keyFile) != 0 {
		log.Info().Str("key_file", keyFile).Str("cert_file", certFile).Msg("running with TLS")

		err := e.RunTLS(address, certFile, keyFile)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to serve")
		}
	} else {
		err := e.Run(address)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to serve")
		}
	}

}
