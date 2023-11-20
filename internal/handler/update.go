package handler

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) handleUpdate(update tgbotapi.Update) error {
	if update.Message == nil {
		slog.Error("Unknown update type", "update", update)

		return nil
	}

	return h.handleMessage(update.Message)
}
