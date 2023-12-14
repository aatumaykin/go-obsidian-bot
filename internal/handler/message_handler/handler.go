package message_handler

import (
	"github.com/aatumaykin/go-obsidian-bot/internal/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHandler interface {
	Handle(note entity.Note, message tgbotapi.Message) (entity.Note, error)
}
