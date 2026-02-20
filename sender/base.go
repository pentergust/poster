package sender

import (
	"poster/parser"
)

// Отправитель сообщения в различные сервисы
type Sender interface {
	Send(message *parser.Message) error
}
