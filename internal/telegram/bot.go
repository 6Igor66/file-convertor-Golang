package telegram

import (
	convertioapi "tgbot/ConvertioAPI"
	"tgbot/internal/storage"
	"time"

	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	tgbot   *tele.Bot
	storage storage.Storage
	api     convertioapi.Api
}

func NewBot(token string, timeout time.Duration, storage storage.Storage, api convertioapi.Api) (*Bot, error) {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: timeout * time.Second},
	}

	var bot Bot
	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}
	bot.tgbot = b
	bot.storage = storage
	bot.api = api
	bot.addHandlers()
	bot.addCallbacks()
	return &bot, nil
}

func (b *Bot) StartBot() {
	b.tgbot.Start()
}
