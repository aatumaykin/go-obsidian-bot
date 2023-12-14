package reply_handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot *tgbotapi.BotAPI
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

func WithBot(bot *tgbotapi.BotAPI) Configuration {
	return func(h *Handler) error {
		h.bot = bot

		return nil
	}
}

func (h *Handler) Handle(update tgbotapi.Update) error {
	mc := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: update.FromChat().ID,
		},
		Text: "...❌... unsupported message",
	}

	if _, err := h.bot.Request(mc); err != nil {
		return err
	}

	return nil
}
