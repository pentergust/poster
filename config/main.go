package config

import (
	"os"

	"github.com/pelletier/go-toml"
)

// Настройки постера
// Предоставляет данные для отправки в разных сервисы
type Config struct {
	Telegram *TelegramConfig
	Discord  *DiscordConfig
	Senders  *[]string
}

// Telegram config
// ===============

// Данные для авторизации Telegram бота
type TelegramBot struct {
	Token string
}

type TelegramChannel struct {
	Bot string
	Id  int64
}

// Настройка отправки сообщений Telegram
// Сообщения в канал отправляются через ботов
// Потому для отправки надо указывать в какой канал какой бот будет отправлять
// сообщения.
type TelegramConfig struct {
	Bot     map[string]TelegramBot
	Channel map[string]TelegramChannel
}

// Настройки Discord
// =================

// Webhook, который отправляет сообщение в канал Discord
type DiscordWebhook struct {
	Username  *string
	AvatarURL *string
	Url       string
	ThreadID  *int
}

// Настройки отправки сообщений в Discord
// Сообщения в канал отправляются через webhook ссылку.
type DiscordConfig struct {
	Webhook map[string]DiscordWebhook
}

// Загружает настройки из файла
func LoadFile(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = toml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
