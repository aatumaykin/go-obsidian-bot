package notify

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Notify struct {
	Bot *tgbotapi.BotAPI
}

func (n *Notify) Reply(chatID int64, messageID int) error {
	mc := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: messageID,
		},
		Text: "...✅...",
	}

	if _, err := n.Bot.Request(mc); err != nil {
		return err
	}

	return nil
}

func (n *Notify) Error(chatID int64, messageID int, msgError string) error {
	mc := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: messageID,
		},
		Text: "...❌..." + msgError,
	}

	if _, err := n.Bot.Request(mc); err != nil {
		return err
	}

	return nil
}
