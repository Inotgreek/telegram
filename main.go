package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5046833680:AAH7WEynACc1KM8Td6aktuUC46RKSvkrtO4")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			command := strings.ToUpper(update.Message.Text)

			switch command {
			case "MTS":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ищу самые большие скидки в МТС")
				bot.Send(msg)
			case "SVY":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ищу самые большие скидки в Связном")
				bot.Send(msg)
			case "BEE":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ищу самые большие скидки в Билайн")
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неверный формат сообщения. \nМТС=MTS Связной=SVY Билайн=BEE")
				bot.Send(msg)

			}

		}
	}
}
