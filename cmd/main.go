package main

import (
	"github.com/404th/anonymous-letter/cmd/bot"
	"github.com/404th/anonymous-letter/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// 0: loading config
	cfg := config.Load()

	// 1: loading logging
	lg := logrus.New()

	switch cfg.Environment {
	case config.DebugMode:
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	// 2: connecting to telegram bot
	if err := bot.NewBot(cfg, lg); err != nil {
		lg.Panic(err)
		return
	}
}
