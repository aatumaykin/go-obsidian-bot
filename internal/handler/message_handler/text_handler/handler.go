package text_handler

import (
	"fmt"
	"unicode/utf16"

	"github.com/aatumaykin/go-obsidian-bot/internal/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct{}

func NewHandler() (*Handler, error) {
	return &Handler{}, nil
}

func (h *Handler) Handle(note entity.Note, message tgbotapi.Message) (entity.Note, error) {
	text := h.ToMarkdown(message.Text, message.Entities)
	caption := h.ToMarkdown(message.Caption, message.CaptionEntities)

	note.Text = fmt.Sprintf("%s\n\n%s", caption, text)

	return note, nil
}

func (h *Handler) ToMarkdown(text string, entities []tgbotapi.MessageEntity) string {
	insertions := make(map[int]string)

	entities = h.filterEntities(entities)

	for i := len(entities) - 1; i >= 0; i-- {
		e := entities[i]

		var before, after string

		// this is supported by the current markdown
		switch e.Type {
		case "bold":
			before = "**"
			after = "**"
		case "italic":
			before = "*"
			after = "*"
		case "strikethrough":
			before = "~~"
			after = "~~"
		case "code":
			before = "`"
			after = "`"
		case "pre":
			before = "```" + e.Language + "\n"
			after = "```"
		case "text_link":
			before = "["
			after = fmt.Sprintf(`](%s)`, e.URL)
		}

		if before != "" {
			insertions[e.Offset] += before
			insertions[e.Offset+e.Length] = after + insertions[e.Offset+e.Length]
		}
	}

	input := []rune(text)
	var output []rune
	utf16pos := 0

	for i := 0; i < len(input); i++ {
		output = append(output, []rune(insertions[utf16pos])...)
		output = append(output, input[i])
		utf16pos += len(utf16.Encode([]rune{input[i]}))
	}
	output = append(output, []rune(insertions[utf16pos])...)

	return string(output)
}

func (h *Handler) filterEntities(entities []tgbotapi.MessageEntity) []tgbotapi.MessageEntity {
	filtered := make([]tgbotapi.MessageEntity, 0)

	for _, e := range entities {
		switch e.Type {
		case "url":
			continue
		case "text_link":
			if e.Length < 3 {
				continue
			}
		}

		filtered = append(filtered, e)
	}

	return filtered
}
