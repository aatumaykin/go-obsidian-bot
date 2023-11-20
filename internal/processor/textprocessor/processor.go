package textprocessor

import (
	"fmt"
	"unicode/utf16"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Processing(message *tgbotapi.Message) string {
	text := getText(message)
	entities := getEntities(message)

	return messageToMarkdown(text, entities)
}

func getText(message *tgbotapi.Message) string {
	if message.Text != "" {
		return message.Text
	} else if message.Caption != "" {
		return message.Caption
	}

	return ""
}

func getEntities(message *tgbotapi.Message) []tgbotapi.MessageEntity {
	if message.Entities != nil {
		return message.Entities
	} else if message.CaptionEntities != nil {
		return message.CaptionEntities
	}

	return nil
}

func messageToMarkdown(text string, entities []tgbotapi.MessageEntity) string {
	insertions := make(map[int]string)

	for i := len(entities) - 1; i >= 0; i-- {
		entity := entities[i]

		var before, after string

		// this is supported by the current markdown
		switch entity.Type {
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
			before = "```" + entity.Language + "\n"
			after = "```"
		case "text_link":
			before = "["
			after = fmt.Sprintf(`](%s)`, entity.URL)
		}

		if before != "" {
			insertions[entity.Offset] += before
			insertions[entity.Offset+entity.Length] = after + insertions[entity.Offset+entity.Length]
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
