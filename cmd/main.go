package main

import (
	"context"
	"fmt"
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

			switch command[0] {
			case "Рандом", "рандом", "случайный":
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Хорошо, сейчас поищу"))
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, kinopoisk_dev.Random()))
			case "Найди", "Найти", "найди", "найти":
				if len(command) < 2 {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Укажите название фильма вторым параметром"))
					break
				}
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Ищу по названию \"%s\", пожождите пожалуйста", command[1])))
				slice := kinopoisk_dev.FindMovieByName(command[1])
				for i := 0; i < len(slice); i++ {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, slice[i]))
				}
			default:
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Не удалось распознать команду"))
			}
		}
	}
}
