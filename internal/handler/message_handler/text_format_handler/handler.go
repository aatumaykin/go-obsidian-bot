package text_format_handler

import (
	"strings"

	"github.com/aatumaykin/go-obsidian-bot/internal/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct{}

func NewHandler() (*Handler, error) {
	return &Handler{}, nil
}

func (h *Handler) Handle(note entity.Note, _ tgbotapi.Message) (entity.Note, error) {
	var result string

	lines := strings.Split(note.Text, "\n")
	isFirstLine := true

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// ignore empty lines
		if line == "" {
			continue
		}

		if isFirstLine {
			if line[0] != '-' && line[0] != '#' && line[0] != '*' {
				line = "# " + line
			}
		}

		isFirstLine = false

		if line[0] != '-' {
			line += "\n\n"
		}

		result += line
	}

	note.Text = result
	return note, nil
}
