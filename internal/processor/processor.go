package processor

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Processor interface {
	Process(update tgbotapi.Update) error
}
