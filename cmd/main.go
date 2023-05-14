package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Rekunch/films-library/internal/config"
	"github.com/Rekunch/films-library/internal/kinopoisk_dev"
	"github.com/Rekunch/films-library/internal/telegram"
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
				telegram.SendMessage(bot, update.Message.Chat.ID, "Хорошо, сейчас поищу")

				result := kinopoisk_dev.Random()
				telegram.SendMessageWithHtml(bot, update.Message.Chat.ID, result)
			case "найти", "найди":
				if len(command) < 2 {
					telegram.SendMessage(bot, update.Message.Chat.ID, "Укажите название фильма вторым параметром")
					return
				}
				searchResult := kinopoisk_dev.FindMovieByName(command[1])
				if len(searchResult) == 0 {
					telegram.SendMessage(bot, update.Message.Chat.ID, "Не удалось найти фильм")
					return
				}
				for _, result := range searchResult {
					telegram.SendMessageWithHtml(bot, update.Message.Chat.ID, result)
				}
			default:
				telegram.SendMessage(bot, update.Message.Chat.ID, "Не удалось распознать команду")
			}
		}
	}
}
