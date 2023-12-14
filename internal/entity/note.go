package entity

import (
	"time"
)

type Note struct {
	ID   string
	Text string
}

func NewNote() *Note {
	return &Note{
		ID: createID(),
	}
}

func createID() string {
	currentTime := time.Now()
	timestamp := currentTime.Format("20060102150405")

	return timestamp
}
