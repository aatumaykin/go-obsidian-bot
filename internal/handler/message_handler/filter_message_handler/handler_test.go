package filter_message_handler_test

import (
	"fmt"
	"testing"

	"github.com/aatumaykin/go-obsidian-bot/internal/entity"
	"github.com/aatumaykin/go-obsidian-bot/internal/handler/message_handler/filter_message_handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Handle(t *testing.T) {
	tests := []struct {
		name    string
		userID  int64
		note    entity.Note
		message tgbotapi.Message
		wantErr error
	}{
		{
			name:   "Message from another user",
			userID: 1,
			note:   entity.Note{},
			message: tgbotapi.Message{
				From: &tgbotapi.User{
					ID: 2,
				},
			},
			wantErr: fmt.Errorf("%w: %s", filter_message_handler.ErrFilterMessage, "message from another user"),
		},
		{
			name:   "Message text is empty",
			userID: 1,
			note:   entity.Note{},
			message: tgbotapi.Message{
				From: &tgbotapi.User{
					ID: 1,
				},
			},
			wantErr: fmt.Errorf("%w: %s", filter_message_handler.ErrFilterMessage, "message text is empty"),
		},
		{
			name:   "Message text is not empty",
			userID: 1,
			note:   entity.Note{},
			message: tgbotapi.Message{
				From: &tgbotapi.User{
					ID: 1,
				},
				Text: "text",
			},
		},
		{
			name:   "Message caption is not empty",
			userID: 1,
			note:   entity.Note{},
			message: tgbotapi.Message{
				From: &tgbotapi.User{
					ID: 1,
				},
				Caption: "text",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, err := filter_message_handler.NewHandler(
				filter_message_handler.WithUserID(tt.userID),
			)
			assert.NoError(t, err)

			note, err := h.Handle(tt.note, tt.message)

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.note, note)
			}
		})
	}
}
