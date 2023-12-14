package helper

import (
	"os"
)

const perm0644 = 0644
const perm0755 = 0755

func CreateFolderIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, perm0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateNote(fileName string, text string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the text to the file
	_, err = file.WriteString(text)
	if err != nil {
		return err
	}

	return nil
}
