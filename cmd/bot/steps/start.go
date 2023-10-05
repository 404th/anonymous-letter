package steps

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var StartKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🏘️ Create a new group", "New group created ✅"),
		tgbotapi.NewInlineKeyboardButtonData("➡️ Join group", "Enter group id ⬇️"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("👥 List of groups", "List of groups 🟰"),
	),
)

var ListOfGroupsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
)
