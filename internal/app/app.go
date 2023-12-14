package app

import (
	"context"
	"log/slog"

	"github.com/aatumaykin/go-obsidian-bot/internal/config"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler/filter_message_handler"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler/remove_handler"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler/reply_handler"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler/save_handler"
	unsupportedReplyHandler "github.com/aatumaykin/go-obsidian-bot/internal/handler/unsupported_handler/reply_handler"
	"github.com/aatumaykin/go-obsidian-bot/internal/processor"
	"github.com/aatumaykin/go-obsidian-bot/internal/processor/message_processor"
	"github.com/aatumaykin/go-obsidian-bot/internal/processor/unsupported_processor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct {
	Context      context.Context
	config       *config.Config
	bot          *tgbotapi.BotAPI
	updateConfig tgbotapi.UpdateConfig
	processors   []processor.Processor
}

type Configuration func(a *App) error

func NewApp(configFile string) (*App, error) {
	a, err := newApp(
		withConfig(configFile),
		withContext(context.Background()),
		withBot(),
		withUpdateConfig(),
		withMessageProcessor(),
		withUnsupportedProcessor(),
	)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func newApp(cfgs ...Configuration) (*App, error) {
	a := &App{}

	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		err := cfg(a)
		if err != nil {
			return nil, err
		}
	}

	return a, nil
}

func withConfig(configFile string) Configuration {
	return func(a *App) error {
		c, err := config.LoadConfig(configFile)
		if err != nil {
			return err
		}

		slog.Info("Config loaded")

		a.config = c

		return nil
	}
}

func withContext(ctx context.Context) Configuration {
	return func(a *App) error {
		a.Context = ctx
		return nil
	}
}

func withBot() Configuration {
	return func(a *App) error {
		bot, err := tgbotapi.NewBotAPI(a.config.Telegram.BotToken)
		if err != nil {
			return err
		}

		a.bot = bot

		return nil
	}
}

func withUpdateConfig() Configuration {
	return func(a *App) error {
		updateConfig := tgbotapi.NewUpdate(0)
		updateConfig.Timeout = a.config.Telegram.Timeout

		a.updateConfig = updateConfig

		return nil
	}
}

func withMessageProcessor() Configuration {
	return func(a *App) error {
		messageProcessor, err := message_processor.NewProcessor(
			message_processor.WithFilterMessageHandler(
				filter_message_handler.WithUserID(a.config.Telegram.UserID),
			),
			message_processor.WithTextHandler(),
			message_processor.WithTextFormatHandler(),
			message_processor.WithSaveHandler(
				save_handler.WithRootPath(a.config.Obsidian.Root),
				save_handler.WithNotePath(a.config.Obsidian.NotePath),
			),
			message_processor.WithReplyHandler(
				reply_handler.WithBot(a.bot),
				reply_handler.WithNeedReply(a.config.Telegram.NeedReply),
			),
			message_processor.WithRemoveHandler(
				remove_handler.WithNeedRemove(a.config.Telegram.NeedRemove),
				remove_handler.WithBot(a.bot),
			),
		)
		if err != nil {
			return err
		}

		a.processors = append(a.processors, messageProcessor)

		return nil
	}
}

func withUnsupportedProcessor() Configuration {
	return func(a *App) error {
		p, err := unsupported_processor.NewProcessor(
			unsupported_processor.WithReplyHandler(
				unsupportedReplyHandler.WithBot(a.bot),
			),
		)
		if err != nil {
			return err
		}

		a.processors = append(a.processors, p)

		return nil
	}
}

func (a *App) Run() error {
	updates := a.bot.GetUpdatesChan(a.updateConfig)

	for update := range updates {
		a.process(update)
	}
	return nil
}

func (a *App) NotifyRun() error {
	slog.Info("Бот Запущен")

	msg := tgbotapi.NewMessage(a.config.Telegram.UserID, "Бот Запущен")
	if _, err := a.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (a *App) process(update tgbotapi.Update) {
	for _, p := range a.processors {
		err := p.Process(update)
		if err != nil {
			slog.Error("Failed to process update", "error", err)
		}
	}
}
