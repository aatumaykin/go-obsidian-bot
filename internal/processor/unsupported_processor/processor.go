package unsupported_processor

import (
	"log/slog"

	"github.com/aatumaykin/go-obsidian-bot/internal/handler/unsupported_handler"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/unsupported_handler/reply_handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Processor struct {
	handlers []unsupported_handler.UnsupportedHandler
}

type Configuration func(p *Processor) error

func NewProcessor(cfgs ...Configuration) (*Processor, error) {
	p := &Processor{}

	for _, cfg := range cfgs {
		err := cfg(p)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func WithReplyHandler(cfgs ...reply_handler.Configuration) Configuration {
	return func(p *Processor) error {
		h, err := reply_handler.NewHandler(cfgs...)
		if err != nil {
			return err
		}

		p.handlers = append(p.handlers, h)

		return nil
	}
}

func (u *Processor) Process(update tgbotapi.Update) error {
	if update.Message != nil {
		return nil
	}

	slog.Warn("Getting unsupported type", "update", update)

	return nil
}
