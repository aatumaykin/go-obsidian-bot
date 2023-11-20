package main

import (
	"log/slog"

	"git.home/Telegram_Bot/go-obsidian-bot/internal/app"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			slog.Error(err.(error).Error())
		}
	}()

	a, err := app.New()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	if err := a.Run(); err != nil {
		slog.Error(err.Error())
	}
}
