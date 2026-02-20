package sender

import (
	"encoding/json"
	"fmt"
	"net/http"
	"poster/parser"
	"time"
)

type TelegramSender struct {
	Token  string
	ChatID int64
}

type telegramMessage struct {
	ChatID    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func (s TelegramSender) toTelegram(text string) telegramMessage {
	return telegramMessage{
		ChatID: s.ChatID,
		Text:   text,
		ParseMode: "html",
	}
}

func (s TelegramSender) Send(message *parser.Message) error {
	tm := s.toTelegram(parser.ToTelegram((*message.Content)))
	payload, err := json.Marshal(tm)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", s.Token)
	return SendRequest(client, url, payload)
}
