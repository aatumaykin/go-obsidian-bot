package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Telegram struct {
	BotToken   string `yaml:"bot_token"`
	Timeout    int    `yaml:"bot_timeout" env-default:"60"`
	UserID     int64  `yaml:"user"`
	NeedRemove bool   `yaml:"need_remove" env-default:"false"`
	NeedReply  bool   `yaml:"need_reply" env-default:"true"`
}

type Obsidian struct {
	Root     string `yaml:"root"`
	NotePath string `yaml:"note_path"`
}

type Config struct {
	Telegram Telegram `yaml:"telegram"`
	Obsidian Obsidian `yaml:"obsidian"`
}

func LoadConfig(configFile string) (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(configFile, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
