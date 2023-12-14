package message_processor

import (
	"github.com/aatumaykin/go-obsidian-bot/internal/entity"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler/filter_message_handler"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler/remove_handler"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler/reply_handler"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler/save_handler"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler/text_format_handler"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler/text_handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Processor struct {
	handlers []message_handler.MessageHandler
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

func WithFilterMessageHandler(cfgs ...filter_message_handler.Configuration) Configuration {
	return func(p *Processor) error {
		h, err := filter_message_handler.NewHandler(cfgs...)
		if err != nil {
			return err
		}

		p.handlers = append(p.handlers, h)

		return nil
	}
}

func WithTextHandler() Configuration {
	return func(p *Processor) error {
		h, err := text_handler.NewHandler()
		if err != nil {
			return err
		}

		p.handlers = append(p.handlers, h)

		return nil
	}
}

func WithTextFormatHandler() Configuration {
	return func(p *Processor) error {
		h, err := text_format_handler.NewHandler()
		if err != nil {
			return err
		}

		p.handlers = append(p.handlers, h)

		return nil
	}
}

func WithSaveHandler(cfgs ...save_handler.Configuration) Configuration {
	return func(p *Processor) error {
		h, err := save_handler.NewHandler(cfgs...)
		if err != nil {
			return err
		}

		p.handlers = append(p.handlers, h)

		return nil
	}
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

func WithRemoveHandler(cfgs ...remove_handler.Configuration) Configuration {
	return func(p *Processor) error {
		h, err := remove_handler.NewHandler(cfgs...)
		if err != nil {
			return err
		}

		p.handlers = append(p.handlers, h)

		return nil
	}
}

func (p *Processor) AddHandler(h message_handler.MessageHandler) {
	p.handlers = append(p.handlers, h)
}

func (p *Processor) Process(update tgbotapi.Update) error {
	if update.Message == nil {
		return nil
	}

	note := *entity.NewNote()
	for _, h := range p.handlers {
		n, err := h.Handle(note, *update.Message)
		if err != nil {
			return err
		}

		note = n
	}
	return nil
}
