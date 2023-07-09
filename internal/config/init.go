package config

import (
	"os"

	"github.com/deka-microservices/go-url-shortener/internal/consts"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	viper.SetEnvPrefix("URL_SHORTENER")
	viper.AutomaticEnv()

	viper.SetDefault(consts.CONFIG_PORT, 9000)
}
