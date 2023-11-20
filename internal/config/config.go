package config

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Telegram struct {
		BotToken   string `yaml:"bot_token"`
		Timeout    int    `yaml:"bot_timeout" env-default:"60"`
		UserID     int64  `yaml:"user"`
		NeedRemove bool   `yaml:"need_remove" env-default:"false"`
		NeedReply  bool   `yaml:"need_reply" env-default:"false"`
	} `yaml:"telegram"`
	Obsidian struct {
		Root     string `yaml:"root"`
		NotePath string `yaml:"note_path"`
	} `yaml:"obsidian"`
}

var ConfigFile = "config.yaml"

func LoadConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(ConfigFile, &cfg)
	if err != nil {
		return nil, err
	}

	slog.Info("Config loaded")

	return &cfg, nil
}
