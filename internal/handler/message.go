package handler

import (
	"git.home/Telegram_Bot/go-obsidian-bot/internal/processor/textprocessor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) handleMessage(message *tgbotapi.Message) error {
	if h.filter(message) {
		return nil
	}

	text := textprocessor.Processing(message)

	if err := h.writeNote(text); err != nil {
		_ = h.Notify.Error(message.Chat.ID, message.MessageID, err.Error())

		return err
	}

	if h.NeedReply {
		if err := h.Notify.Reply(message.Chat.ID, message.MessageID); err != nil {
			return err
		}
	}

	if h.NeedRemove {
		if err := h.deleteMessage(message.Chat.ID, message.MessageID); err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) deleteMessage(chatID int64, messageID int) error {
	if _, err := h.Bot.Request(tgbotapi.NewDeleteMessage(chatID, messageID)); err != nil {
		return err
	}

	return nil
}
