package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func SendMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := bot.Send(msg)
	if err != nil {
		return
	}
}

func SendMessageWithHtml(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = "html"
	_, err := bot.Send(msg)
	if err != nil {
		return
	}
}
