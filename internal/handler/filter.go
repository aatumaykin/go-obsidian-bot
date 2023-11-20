package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handler) filter(message *tgbotapi.Message) bool {
	return h.denyForUser(message.From.ID)
}

func (h *Handler) denyForUser(userID int64) bool {
	return userID != h.UserID
}
