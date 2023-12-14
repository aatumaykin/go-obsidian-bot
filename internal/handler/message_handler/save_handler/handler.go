package save_handler

import (
	"path"

	"github.com/aatumaykin/go-obsidian-bot/internal/entity"
	"github.com/aatumaykin/go-obsidian-bot/internal/helper"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	rootPath string
	notePath string
}

type Configuration func(h *Handler) error

func NewHandler(cfgs ...Configuration) (*Handler, error) {
	h := &Handler{}

	for _, cfg := range cfgs {
		err := cfg(h)
		if err != nil {
			return nil, err
		}
	}

	return h, nil
}

func WithRootPath(rootPath string) Configuration {
	return func(h *Handler) error {
		h.rootPath = rootPath
		return nil
	}
}

func WithNotePath(notePath string) Configuration {
	return func(h *Handler) error {
		h.notePath = notePath
		return nil
	}
}

func (h *Handler) Handle(note entity.Note, _ tgbotapi.Message) (entity.Note, error) {
	if err := helper.CreateFolderIfNotExist(path.Join(h.rootPath, h.notePath)); err != nil {
		return note, err
	}

	fileName := path.Join(h.rootPath, h.notePath, note.ID+".md")

	return note, helper.CreateNote(fileName, note.Text)
}
