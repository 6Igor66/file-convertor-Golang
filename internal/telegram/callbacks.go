package telegram

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"tgbot/internal/service"
	"time"

	tele "gopkg.in/telebot.v3"
)

func (b *Bot) addCallbacks() {
	b.tgbot.Handle(tele.OnCallback, func(c tele.Context) error {
		file, _ := b.storage.GetPayload(c.Sender().ID)
		fileName, err := b.storage.GetFileName(c.Sender().ID)
		if err != nil {
			log.Printf("\nError get filename: %s", err)
			return c.Send("Что-то пошло не так...")
		}
		format := c.Callback().Data
		payload := []string{file, format}
		if fileName != "" {
			payload = append(payload, fileName, "base64")
		}
		content, err := service.TransformFile(b.api, payload)
		if err != nil {
			log.Printf("\nError Transform file: %s", err)
			return c.Send("Что-то пошло не так...")
		}

		pdfData, err := base64.StdEncoding.DecodeString(content)
		if err != nil {
			return c.Send("Не смог задекодить файл")
		}

		pdfBuffer := bytes.NewReader(pdfData)

		var result *tele.Document
		switch format {
		case "pdf":
			result = &tele.Document{
				File:     tele.FromReader(pdfBuffer),
				FileName: "output.pdf",
				MIME:     "application/pdf",
			}
		case "doc":
			result = &tele.Document{
				File:     tele.FromReader(pdfBuffer),
				FileName: "output.doc",
				MIME:     "application/msword",
			}
		case "png":
			result = &tele.Document{
				File:     tele.FromReader(pdfBuffer),
				FileName: "output.png",
				MIME:     "image/png",
			}
		case "jpeg":
			result = &tele.Document{
				File:     tele.FromReader(pdfBuffer),
				FileName: "output.jpeg",
				MIME:     "image/jpeg",
			}
		}
		time.Sleep(5 * time.Second)

		c.Send(result)

		err = b.storage.UpdateMessageStatus(c.Sender().ID, "waiting_payload")
		if err != nil {
			log.Printf("\nError update user status: %s", err)
			return err
		}

		nums, err := b.storage.GetOperations(c.Sender().ID)
		if err != nil {
			log.Printf("\nError get operations user: %s", err)
			return err
		}

		err = b.storage.SetOperations(c.Sender().ID, nums-1)
		if err != nil {
			log.Printf("\nError update operations user: %s", err)
			return err
		}

		err = b.storage.SetFileName(c.Sender().ID, "")
		if err != nil {
			log.Printf("\nError update filename: %s", err)
			return err
		}
		msg := fmt.Sprintf("У вас осталось %d использований\nПришлите URL того, что нужно преобразовать", nums-1)
		return c.Send(msg)
	})
}
