package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/teatou/TelegramBot/telegram"
)

func init() {
	godotenv.Load()
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_KEY"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
