package remove_handler

import (
	"github.com/aatumaykin/go-obsidian-bot/internal/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	needRemove bool
	bot        *tgbotapi.BotAPI
}

type Configuration func(h *Handler) error

func NewHandler(cfgs ...Configuration) (*Handler, error) {
	h := &Handler{}

	for _, cfg := range cfgs {
		err := cfg(h)
		if err != nil {
			return nil, err
		}
	}

	return h, nil
}

func WithNeedRemove(needRemove bool) Configuration {
	return func(h *Handler) error {
		h.needRemove = needRemove

		return nil
	}
}

func WithBot(bot *tgbotapi.BotAPI) Configuration {
	return func(h *Handler) error {
		h.bot = bot

		return nil
	}
}

func (h *Handler) Handle(note entity.Note, message tgbotapi.Message) (entity.Note, error) {
	if !h.needRemove {
		return note, nil
	}

	if _, err := h.bot.Request(tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)); err != nil {
		return note, err
	}

	return note, nil
}
