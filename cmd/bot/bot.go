package bot

import (
	"fmt"

	"github.com/404th/anonymous-letter/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func NewBot(cfg config.Config, lg *logrus.Logger) error {
	bot := &tgbotapi.BotAPI{}

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		lg.Panicf("[%s] [error] [%s]", cfg.Environment, err.Error())
		return err
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			lg.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			str := fmt.Sprintf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, str)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}

	return nil
}
