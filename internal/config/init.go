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
	viper.SetDefault(consts.CONFIG_FAIL_PERCENT, uint32(25))

	viper.SetDefault(consts.CONFIG_REDIS_ADDRESSES, "")
	viper.SetDefault(consts.CONFIG_REDIS_PASSWORD, "")

	viper.SetDefault(consts.CONFIG_USE_TLS, false)
	viper.SetDefault(consts.CONFIG_TLS_CERT, "")
	viper.SetDefault(consts.CONFIG_TLS_KEY, "")
}
