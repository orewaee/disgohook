package disgohook

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Webhook struct {
	url  string
	body *Body
}

type Body struct {
	Content string `json:"content"`
}

func (webhook *Webhook) Send() error {
	body, err := json.Marshal(webhook.body)
	if err != nil {
		return err
	}

	buffer := bytes.NewBuffer(body)
	req, err := http.NewRequest("POST", webhook.url, buffer)
	if err != nil {
		return err
	}

	contentType := "application/json; charset=UTF-8"
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			return
		}
	}()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	ok := res.StatusCode == http.StatusOK ||
		res.StatusCode == http.StatusNoContent

	if !ok {
		return errors.New("failed to send webhook: " + string(bytes))
	}

	return nil
}
