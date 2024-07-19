package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Bot        `yaml:"bot"`
	PostgreSQL `yaml:"postgresql"`
	Http       `yaml:"http"`
}

type Bot struct {
	BotToken string        `yaml:"bot_token"`
	Timeout  time.Duration `yaml:"poller_timeout_ms"`
	Debug    bool          `yaml:"debug"`
}

type PostgreSQL struct {
	Connstring string `yaml:"connstring"`
}

type Http struct {
	ApiKey string `yaml:"api_key"`
}

func MustLoad(configPath string) (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
