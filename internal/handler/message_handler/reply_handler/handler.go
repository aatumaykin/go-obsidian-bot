package reply_handler

import (
	"github.com/aatumaykin/go-obsidian-bot/internal/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	needReply bool
	bot       *tgbotapi.BotAPI
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

func WithNeedReply(needReply bool) Configuration {
	return func(h *Handler) error {
		h.needReply = needReply

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
	if !h.needReply {
		return note, nil
	}

	mc := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           message.Chat.ID,
			ReplyToMessageID: message.MessageID,
		},
		Text: "...✅...",
	}

	if _, err := h.bot.Request(mc); err != nil {
		return note, err
	}

	return note, nil
}
