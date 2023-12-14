package message_processor_test

import (
	"fmt"
	"testing"

	"github.com/aatumaykin/go-obsidian-bot/internal/entity"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler/filter_message_handler"
	"github.com/aatumaykin/go-obsidian-bot/internal/processor/message_processor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

type TestHandler struct {
	Note entity.Note
}

func (h *TestHandler) Handle(note entity.Note, _ tgbotapi.Message) (entity.Note, error) {
	h.Note = note

	return note, nil
}

func TestMessageProcessor_Process(t *testing.T) {
	tests := []struct {
		name     string
		userID   int64
		update   tgbotapi.Update
		wantText string
		wantErr  error
	}{
		{
			name:   "Test filter user",
			userID: 1,
			update: tgbotapi.Update{
				Message: &tgbotapi.Message{
					From: &tgbotapi.User{ID: 2},
				},
			},
			wantErr: fmt.Errorf("%w: %s", filter_message_handler.ErrFilterMessage, "message from another user"),
		},
		{
			name:   "Test filter text",
			userID: 1,
			update: tgbotapi.Update{
				Message: &tgbotapi.Message{
					From: &tgbotapi.User{ID: 1},
				},
			},
			wantErr: fmt.Errorf("%w: %s", filter_message_handler.ErrFilterMessage, "message text is empty"),
		},
		{
			// в этом канале любят вставлять ссылки на канал в точки и запятые
			name:   "Test filter text format",
			userID: 1,
			update: tgbotapi.Update{
				Message: &tgbotapi.Message{
					From: &tgbotapi.User{ID: 1},
					Caption: `Oatmeal

Oatmeal - это чат-приложение с терминальным пользовательским интерфейсом, которое взаимодействует с большими языковыми моделями (LLM), используя различные бэкенды моделей.

☝🏻Он может быть интегрирован с такими редакторами, как Neovim.

Oatmeal позволяет настраивать конфигурацию с помощью параметров командной строки или переменных окружения, поддерживает слэш-команды и пузырьки чата.

Приложение обеспечивает подсветку синтаксиса кода с различными темами.

Все сеансы чата с моделями сохраняются, что позволяет пользователям просматривать или продолжать предыдущие разговоры.

https://github.com/dustinblackman/oatmeal

Site: https://dustinblackman.com/posts/oatmeal/`,
					CaptionEntities: []tgbotapi.MessageEntity{
						tgbotapi.MessageEntity{
							Type:   "bold",
							Offset: 0,
							Length: 7,
						},
						tgbotapi.MessageEntity{
							Type:   "text_link",
							Offset: 81,
							Length: 1,
							URL:    "https://t.me/open_source_friend",
						},
						tgbotapi.MessageEntity{
							Type:   "text_link",
							Offset: 142,
							Length: 1,
							URL:    "https://t.me/open_source_friend",
						},
						tgbotapi.MessageEntity{
							Type:   "text_link",
							Offset: 179,
							Length: 1,
							URL:    "https://t.me/open_source_friend",
						},
						tgbotapi.MessageEntity{
							Type:   "text_link",
							Offset: 232,
							Length: 1,
							URL:    "https://t.me/open_source_friend",
						},
						tgbotapi.MessageEntity{
							Type:   "text_link",
							Offset: 244,
							Length: 1,
							URL:    "https://t.me/open_source_friend",
						},
						tgbotapi.MessageEntity{
							Type:   "text_link",
							Offset: 352,
							Length: 1,
							URL:    "https://t.me/open_source_friend",
						},
						tgbotapi.MessageEntity{
							Type:   "text_link",
							Offset: 395,
							Length: 1,
							URL:    "https://t.me/open_source_friend",
						},
						tgbotapi.MessageEntity{
							Type:   "text_link",
							Offset: 467,
							Length: 1,
							URL:    "https://t.me/open_source_friend",
						},
						tgbotapi.MessageEntity{
							Type:   "text_link",
							Offset: 508,
							Length: 1,
							URL:    "https://t.me/open_source_friend",
						},
						tgbotapi.MessageEntity{
							Type:   "text_link",
							Offset: 587,
							Length: 1,
							URL:    "https://t.me/open_source_friend",
						},
						tgbotapi.MessageEntity{
							Type:   "url",
							Offset: 590,
							Length: 41,
						},
						tgbotapi.MessageEntity{
							Type:   "url",
							Offset: 639,
							Length: 41,
						},
					},
				},
			},
			wantText: `**Oatmeal**

Oatmeal - это чат-приложение с терминальным пользовательским интерфейсом, которое взаимодействует с большими языковыми моделями (LLM), используя различные бэкенды моделей.

☝🏻Он может быть интегрирован с такими редакторами, как Neovim.

Oatmeal позволяет настраивать конфигурацию с помощью параметров командной строки или переменных окружения, поддерживает слэш-команды и пузырьки чата.

Приложение обеспечивает подсветку синтаксиса кода с различными темами.

Все сеансы чата с моделями сохраняются, что позволяет пользователям просматривать или продолжать предыдущие разговоры.

https://github.com/dustinblackman/oatmeal

Site: https://dustinblackman.com/posts/oatmeal/

`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testHandler := &TestHandler{}

			p, err := message_processor.NewProcessor(
				message_processor.WithFilterMessageHandler(
					filter_message_handler.WithUserID(tt.userID),
				),
				message_processor.WithTextHandler(),
				message_processor.WithTextFormatHandler(),
			)
			assert.NoError(t, err)

			p.AddHandler(testHandler)
			err = p.Process(tt.update)

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantText, testHandler.Note.Text)
			}
		})
	}
}
