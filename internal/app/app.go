package app

import (
	"log/slog"
	"runtime"

	"git.home/Telegram_Bot/go-obsidian-bot/internal/config"
	"git.home/Telegram_Bot/go-obsidian-bot/internal/handler"
	"git.home/Telegram_Bot/go-obsidian-bot/internal/notify"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct {
	Config       *config.Config
	Bot          *tgbotapi.BotAPI
	UpdateConfig tgbotapi.UpdateConfig
}

func New() (*App, error) {
	slog.Info("Logger created")
	slog.Info("Go version", "version", runtime.Version())

	c, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	bot, err := tgbotapi.NewBotAPI(c.Telegram.BotToken)
	if err != nil {
		return nil, err
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = c.Telegram.Timeout

	a := &App{
		Config:       c,
		Bot:          bot,
		UpdateConfig: updateConfig,
	}

	return a, nil
}

func (a *App) Run() error {
	if err := a.notifyRun(); err != nil {
		return err
	}

	updates := a.Bot.GetUpdatesChan(a.UpdateConfig)

	h := handler.Handler{
		Bot:              a.Bot,
		Notify:           &notify.Notify{Bot: a.Bot},
		NeedRemove:       a.Config.Telegram.NeedRemove,
		NeedReply:        a.Config.Telegram.NeedReply,
		UserID:           a.Config.Telegram.UserID,
		ObsidianRootPath: a.Config.Obsidian.Root,
		NotePath:         a.Config.Obsidian.NotePath,
	}
	if err := h.Run(updates); err != nil {
		slog.Error(err.Error())
	}

	return nil
}

func (a *App) notifyRun() error {
	slog.Info("Бот Запущен")

	msg := tgbotapi.NewMessage(a.Config.Telegram.UserID, "Бот Запущен")
	if _, err := a.Bot.Send(msg); err != nil {
		return err
	}

	return nil
}
