package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Rekunch/films-library/internal/config"
	"github.com/Rekunch/films-library/internal/kinopoisk_dev"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	ctx := context.Background()

	err := config.Init(ctx)
	if err != nil {
		log.Panic(err)
	}

	TelegramToken := os.Getenv("TELEGRAM_TOKEN")

	bot, err := tgbotapi.NewBotAPI(TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	timeout, err := strconv.Atoi(os.Getenv("TELEGRAM_TIMEOUT"))
	if err != nil {
		panic(err)
	}
	u.Timeout = timeout

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message

			command := strings.Split(update.Message.Text, " ")

			switch strings.ToLower(command[0]) {
			case "рандом", "случайный":
				sendMessage(bot, update.Message.Chat.ID, "Хорошо, сейчас поищу")

				result := kinopoisk_dev.Random()
				sendMessageWithHtml(bot, update.Message.Chat.ID, result)
			case "найти", "найди":
				if len(command) < 2 {
					sendMessage(bot, update.Message.Chat.ID, "Укажите название фильма вторым параметром")
					return
				}
				searchResult := kinopoisk_dev.FindMovieByName(command[1])
				if len(searchResult) == 0 {
					sendMessage(bot, update.Message.Chat.ID, "Не удалось найти фильм")
					return
				}
				for _, result := range searchResult {
					sendMessageWithHtml(bot, update.Message.Chat.ID, result)
				}
			default:
				sendMessage(bot, update.Message.Chat.ID, "Не удалось распознать команду")
			}
		}
	}
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := bot.Send(msg)
	if err != nil {
		return
	}
}

func sendMessageWithHtml(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = "html"
	_, err := bot.Send(msg)
	if err != nil {
		return
	}
}
