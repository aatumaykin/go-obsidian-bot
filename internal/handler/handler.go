package handler

import (
	"git.home/Telegram_Bot/go-obsidian-bot/internal/notify"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	Bot              *tgbotapi.BotAPI
	Notify           *notify.Notify
	NeedRemove       bool
	NeedReply        bool
	UserID           int64
	ObsidianRootPath string
	NotePath         string
}

func (h *Handler) Run(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if err := h.handleUpdate(update); err != nil {
			return err
		}
	}

	return nil
}
