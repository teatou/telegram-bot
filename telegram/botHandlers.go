package telegram

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/teatou/telegram-bot/dict"
)

const (
	commandStart       = "start"
	commandHelp        = "help"
	commandReset       = "reset"
	commandChangeLangs = "change"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command")
	switch message.Command() {
	case commandStart:
		msg.Text = "Write any word you want to study or translate. Write /help for info"
		_, err := b.bot.Send(msg)
		return err
	case commandHelp:
		msg.Text = "To change languages, write '/change en-en' with such options: en, ru, es, it, de, fr.\nIf something went wrong write /reset"
		_, err := b.bot.Send(msg)
		return err
	case commandReset:
		b.langs = []string{"en", "en"}
		msg.Text = "Languages have been reset to en-en"
		_, err := b.bot.Send(msg)
		return err
	case commandChangeLangs:
		cmd := strings.Split(message.Text, " ")
		if len(cmd) == 1 {
			msg.Text = "Failed to change languages. Write '/change en-en' with such options: en, ru, es, it, de, fr"
		} else {
			b.langs = strings.Split(cmd[1], "-")
			msg.Text = "Languages changed"
		}
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
	msgText, err := dict.Lookup(lookup, b.langs)
	if err != nil || msgText == "" {
		msg = tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Unable to handle request. Current languages: %s-%s. Write /help for more info", b.langs[0], b.langs[1]))
		log.Println("error looking up dictionary")
	} else {
		msg = tgbotapi.NewMessage(message.Chat.ID, msgText)
	}

	b.bot.Send(msg)
}
