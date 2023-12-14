package filter_message_handler

import (
	"errors"
	"fmt"

	"github.com/aatumaykin/go-obsidian-bot/internal/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	userID int64
}

type Configuration func(h *Handler) error

var ErrFilterMessage = errors.New("filter message error")

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

func WithUserID(userID int64) Configuration {
	return func(h *Handler) error {
		h.userID = userID
		return nil
	}
}

func (h *Handler) Handle(note entity.Note, message tgbotapi.Message) (entity.Note, error) {
	if message.From.ID != h.userID {
		return note, fmt.Errorf("%w: %s", ErrFilterMessage, "message from another user")
	}

	if message.Text == "" && message.Caption == "" {
		return note, fmt.Errorf("%w: %s", ErrFilterMessage, "message text is empty")
	}

	return note, nil
}
