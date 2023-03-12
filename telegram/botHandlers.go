package telegram

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/teatou/telegram-bot/dict"
)

const (
	commandStart = "start"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда")
	switch message.Command() {
	case commandStart:
		msg.Text = "command start"
		_, err := b.bot.Send(msg)
		return err

	default:
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	var msg tgbotapi.MessageConfig

	lookup := strings.Split(message.Text, " ")
	msgText, err := dict.Lookup(lookup)
	if err != nil || msgText == "" {
		msg = tgbotapi.NewMessage(message.Chat.ID, "Не удалось обработать запрос")
		log.Println("error looking up in dictionary")
	} else {
		msg = tgbotapi.NewMessage(message.Chat.ID, msgText)
	}

	b.bot.Send(msg)
}
