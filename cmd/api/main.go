package main

import (
	"fmt"
	"log"
	"net/http"

	// _ "order-service/docs"
	"order-service/internal/app/config"
	"order-service/internal/app/db"
	"order-service/internal/app/router"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func NewGinServer(config *config.Configuration) *http.Server {
	r := router.Setup(config)
	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", config.Server.Port),
	}
	return s
}

func setConfiguration(configPath string) {
	config.Setup(configPath)
	db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func main() {
	setConfiguration("configs/config-postgres.yml")
	// setConfiguration("data/config-postgres.yml")
	conf := config.GetConfig()
	// configureLogger(conf)
	s := NewGinServer(conf)
	log.Fatal(s.ListenAndServe())
}

func configureLogger(conf *config.Configuration) {
	switch conf.Logging.Level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)

	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)

	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)

	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "disabled":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	case "":
		zerolog.SetGlobalLevel(zerolog.NoLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)

	}
}
