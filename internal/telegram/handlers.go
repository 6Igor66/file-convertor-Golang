package telegram

import (
	"encoding/base64"
	"io"
	"log"

	tele "gopkg.in/telebot.v3"
)

func (b *Bot) addHandlers() {
	b.tgbot.Handle("/start", func(c tele.Context) error {
		b.storage.CreateUser(c.Sender().ID)
		return c.Send("Привет, чтобы продолжить, отправь файл или URL")
	})
	b.tgbot.Handle(tele.OnText, func(c tele.Context) error {
		status, err := b.storage.GetMessageStatus(c.Sender().ID)
		if err != nil {
			log.Printf("Cannot get message status: %s", err)
		}

		var inlineKeys []tele.InlineButton
		formats := getSupportedFormats()
		switch status {
		case "waiting_payload":
			file := c.Message().Text
			b.storage.UpdateMessageStatus(c.Sender().ID, "waiting_result")
			b.storage.SetPaylaod(c.Sender().ID, file) //TODO он отдал нам урл, сделать придумать всю дальнейшую логику
			for _, v := range formats {
				button := CreateButton(v)
				inlineKeys = append(inlineKeys, *button)
			}
			inlineKeyboard := tele.ReplyMarkup{InlineKeyboard: [][]tele.InlineButton{inlineKeys}}
			c.Send("Выбери желаемый формат", &inlineKeyboard)
		default:
			c.Send("Не ожидается новых сообщений")
		}
		return nil
	})

	b.tgbot.Handle(tele.OnDocument, func(c tele.Context) error {
		status, err := b.storage.GetMessageStatus(c.Sender().ID)
		if err != nil {
			log.Printf("Cannot get message status: %s", err)
		}
		switch status {
		case "waiting_payload":
			file := c.Message().Document

			reader, err := b.tgbot.File(file.MediaFile())
			if err != nil {
				return c.Send("Что-то пошло не так")
			}
			defer reader.Close()

			fileBytes, err := io.ReadAll(reader)
			if err != nil {
				return err
			}

			encodedFile := base64.StdEncoding.EncodeToString(fileBytes)
			fileName := file.FileName

			b.storage.SetPaylaod(c.Sender().ID, encodedFile)
			b.storage.UpdateMessageStatus(c.Sender().ID, "waiting_result")
			b.storage.SetFileName(c.Sender().ID, fileName)

			var inlineKeys []tele.InlineButton
			formats := getSupportedFormats()
			for _, v := range formats {
				button := CreateButton(v)
				inlineKeys = append(inlineKeys, *button)
			}
			inlineKeyboard := tele.ReplyMarkup{InlineKeyboard: [][]tele.InlineButton{inlineKeys}}
			c.Send("Выбери желаемый формат", &inlineKeyboard)

		default:
			return c.Send("Не ожидается новых сообщений")
		}
		return nil
	})
}

func CreateButton(data string) *tele.InlineButton {
	button := tele.InlineButton{
		Data: data,
		Text: data,
	}
	return &button
}

func getSupportedFormats() []string {
	return []string{"pdf", "png", "doc", "jpeg"}
}
