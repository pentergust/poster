package parser

import (
	"os"

	"github.com/pelletier/go-toml"
)


type Message struct {
	Content   *string
}


// Загружает настройки из файла
func LoadMessage(path string) (*Message, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var message Message
	err = toml.Unmarshal(file, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
