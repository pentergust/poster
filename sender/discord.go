package sender

import (
	"encoding/json"
	"errors"
	"net/http"
	"poster/parser"
	"time"
)

type DiscordSender struct {
	Username  *string
	AvatarURL *string
	Url       string
	ThreadID  *int
}

type discordMessage struct {
	Content   *string `json:"content,omitempty"`
	Username  *string `json:"username,omitempty"`
	AvatarURL *string `json:"avatar_url,omitempty"`
}

func (s *DiscordSender) toDiscord(message *parser.Message) (*discordMessage, error) {
	if message.Content == nil {
		return nil, errors.New("Message has no content to send")
	}

	return &discordMessage{
		Content:   message.Content,
		Username:  s.Username,
		AvatarURL: s.AvatarURL,
	}, nil
}

func (s DiscordSender) Send(message *parser.Message) error {
	dm, err := s.toDiscord(message)
	if err != nil {
		return err
	}

	payload, err := json.Marshal(dm)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	return SendRequest(client, s.Url, payload)
}
