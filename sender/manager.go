package sender

import (
	"poster/config"
	"poster/parser"

	"github.com/rs/zerolog/log"
)

// Коллекция отправителей
type SenderManager struct {
	Senders map[string]Sender
}

// Запускает отправку сообщений выбранным отправителям
func (m *SenderManager) Send(message *parser.Message, to []string) {
	for _, t := range to {
		sender, ok := m.Senders[t]
		if !ok {
			log.Warn().Str("To", t).Msg("Not found")
			continue
		}

		err := sender.Send(message)
		if err != nil {
			log.Error().Err(err).Send()
		}
	}
}

// Подготавливает всех отправителей на основе настроек
func NewManager(config *config.Config) SenderManager {
	senders := make(map[string]Sender)

	// Telegram bots
	for name, ch := range config.Telegram.Channel {
		bot, ok := config.Telegram.Bot[ch.Bot]
		if !ok {
			log.Warn().Str("Bot", ch.Bot).Str("Chan", name).Msg("Bot not found")
			continue
		}

		senders[name] = TelegramSender{bot.Token, ch.Id}
		log.Debug().Str("Name", name).Int64("Id", ch.Id).Msg("Register telegram channel")
	}

	// Discord Webhooks
	for name, hook := range config.Discord.Webhook {
		senders[name] = DiscordSender{
			Username:  hook.Username,
			AvatarURL: hook.AvatarURL,
			Url:       hook.Url,
			ThreadID:  hook.ThreadID,
		}
	}

	return SenderManager{senders}
}
