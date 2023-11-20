package handler

import (
	"path"

	"git.home/Telegram_Bot/go-obsidian-bot/internal/helper"
)

func (h *Handler) writeNote(text string) error {
	if err := helper.CreateFolderIfNotExist(path.Join(h.ObsidianRootPath, h.NotePath)); err != nil {
		return err
	}

	fileName := path.Join(h.ObsidianRootPath, h.NotePath, helper.GetCurrentTimestampAsString()+".md")
	return helper.CreateNote(fileName, text)
}
