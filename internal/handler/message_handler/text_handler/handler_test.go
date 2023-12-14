package text_handler_test

import (
	"testing"

	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler/text_handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

func TestHandler_ToMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		entities []tgbotapi.MessageEntity
		want     string
	}{
		{
			name:     "Bold",
			text:     "test",
			entities: []tgbotapi.MessageEntity{{Type: "bold", Offset: 0, Length: 4}},
			want:     "**test**",
		},
		{
			name:     "Italic",
			text:     "test",
			entities: []tgbotapi.MessageEntity{{Type: "italic", Offset: 0, Length: 4}},
			want:     "*test*",
		},
		{
			name:     "Strikethrough",
			text:     "test",
			entities: []tgbotapi.MessageEntity{{Type: "strikethrough", Offset: 0, Length: 4}},
			want:     "~~test~~",
		},
		{
			name:     "Code",
			text:     "test",
			entities: []tgbotapi.MessageEntity{{Type: "code", Offset: 0, Length: 4}},
			want:     "`test`",
		},
		{
			name:     "Pre",
			text:     "test",
			entities: []tgbotapi.MessageEntity{{Type: "pre", Offset: 0, Length: 4, Language: "text"}},
			want:     "```text\ntest```",
		},
		{
			name:     "TextLink",
			text:     "test",
			entities: []tgbotapi.MessageEntity{{Type: "text_link", Offset: 0, Length: 4, URL: "test"}},
			want:     "[test](test)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, _ := text_handler.NewHandler()

			result := h.ToMarkdown(tt.text, tt.entities)
			assert.Equal(t, tt.want, result)
		})
	}
}
