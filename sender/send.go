package sender

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

// Вспомогательная функция отправки запроса при помощи клиента
func SendRequest(client *http.Client, url string, payload []byte) error {
	resp, err := client.Post(
		url,
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody := new(bytes.Buffer)
	n, err := respBody.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	log.Info().Int64("Bytes", n).Str("Resp", respBody.String()).Send()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("error response code: %d", resp.StatusCode)
	}
	return nil
}
