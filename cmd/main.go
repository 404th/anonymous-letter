package main

import (
	"log"

	"github.com/404th/anonymous-letter/config"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s\n", bot.Self.UserName)
	log.Printf("Authorized on Firstname account %s\n", &bot.Self.FirstName)
	log.Printf("Authorized on Lastname account %s\n", &bot.Self.LastName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
