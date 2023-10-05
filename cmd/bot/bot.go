package bot

import (
	"fmt"

	"github.com/404th/anonymous-letter/cmd/bot/steps"
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
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Command() {
			case "start":
				fmt.Println(">>> 4")
				msg.ReplyMarkup = steps.StartKeyboard
			default:
				fmt.Println(">>> 5")
				msg.Text = "I don't know that command"
			}

			// switch update.Message.Text {
			// case "List of groups 🟰":
			// 	fmt.Println(">>> 1")
			// 	msg.ReplyMarkup = steps.ListOfGroupsKeyboard
			// default:
			// 	fmt.Println(">>> 2")
			// 	msg.Text = "I don't know that text command"
			// }

			fmt.Println(">>> 3")
			// Send the message.
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}

	return nil
}
