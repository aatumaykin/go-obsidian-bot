package unsupported_handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UnsupportedHandler interface {
	Handle(update tgbotapi.Update) error
}
