package text_format_handler_test

import (
	"testing"

	"github.com/aatumaykin/go-obsidian-bot/internal/entity"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler/text_format_handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Handle(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name: "Text 1",
			text: `**Oatmeal**

Oatmeal - это чат-приложение с терминальным пользовательским интерфейсом[,](https://t.me/open_source_friend) которое взаимодействует с большими языковыми моделями (LLM)[,](https://t.me/open_source_friend) используя различные бэкенды моделей[.](https://t.me/open_source_friend)

☝🏻Он может быть интегрирован с такими редакторами[,](https://t.me/open_source_friend) как Neovim[.](https://t.me/open_source_friend)

Oatmeal позволяет настраивать конфигурацию с помощью параметров командной строки или переменных окружения[,](https://t.me/open_source_friend) поддерживает слэш-команды и пузырьки чата[.](https://t.me/open_source_friend)

Приложение обеспечивает подсветку синтаксиса кода с различными темами[.](https://t.me/open_source_friend)

Все сеансы чата с моделями сохраняются[,](https://t.me/open_source_friend) что позволяет пользователям просматривать или продолжать предыдущие разговоры[.](https://t.me/open_source_friend)

https://github.com/dustinblackman/oatmeal

Site: https://dustinblackman.com/posts/oatmeal/`,
			expected: `**Oatmeal**

Oatmeal - это чат-приложение с терминальным пользовательским интерфейсом[,](https://t.me/open_source_friend) которое взаимодействует с большими языковыми моделями (LLM)[,](https://t.me/open_source_friend) используя различные бэкенды моделей[.](https://t.me/open_source_friend)

☝🏻Он может быть интегрирован с такими редакторами[,](https://t.me/open_source_friend) как Neovim[.](https://t.me/open_source_friend)

Oatmeal позволяет настраивать конфигурацию с помощью параметров командной строки или переменных окружения[,](https://t.me/open_source_friend) поддерживает слэш-команды и пузырьки чата[.](https://t.me/open_source_friend)

Приложение обеспечивает подсветку синтаксиса кода с различными темами[.](https://t.me/open_source_friend)

Все сеансы чата с моделями сохраняются[,](https://t.me/open_source_friend) что позволяет пользователям просматривать или продолжать предыдущие разговоры[.](https://t.me/open_source_friend)

https://github.com/dustinblackman/oatmeal

Site: https://dustinblackman.com/posts/oatmeal/

`,
		},
		{
			name: "Text 2",
			text: `**Oatmeal**

Oatmeal - это чат-приложение с терминальным пользовательским интерфейсом, которое взаимодействует с большими языковыми моделями (LLM), используя различные бэкенды моделей.

☝🏻Он может быть интегрирован с такими редакторами, как Neovim.

Oatmeal позволяет настраивать конфигурацию с помощью параметров командной строки или переменных окружения, поддерживает слэш-команды и пузырьки чата.

Приложение обеспечивает подсветку синтаксиса кода с различными темами.

Все сеансы чата с моделями сохраняются, что позволяет пользователям просматривать или продолжать предыдущие разговоры.

https://github.com/dustinblackman/oatmeal

Site: https://dustinblackman.com/posts/oatmeal/`,
			expected: `**Oatmeal**

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
			h := &text_format_handler.Handler{}

			note := entity.Note{
				Text: tt.text,
			}

			n, err := h.Handle(note, tgbotapi.Message{})
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, n.Text)
		})
	}
}
