package main

import (
	"log"
	convertioapi "tgbot/ConvertioAPI"
	"tgbot/internal/config"
	"tgbot/internal/storage/postgres"
	"tgbot/internal/telegram"
)

func main() {
	cfg, err := config.MustLoad("../../config/local.yaml")
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	storage, err := postgres.NewPostgres(cfg.PostgreSQL.Connstring)
	if err != nil {
		log.Fatalf("error connect db: %s", err)
	}

	api := convertioapi.NewApi(cfg.Http.ApiKey)
	bot, err := telegram.NewBot(cfg.Bot.BotToken, cfg.Bot.Timeout, storage, api)
	if err != nil {
		log.Fatalf("error init bot: %s", err)
	}

	bot.StartBot()

}
