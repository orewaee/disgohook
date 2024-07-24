package disgohook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Webhook struct {
	url  string
	body *Body
}

type Body struct {
	Content string   `json:"content"`
	Embeds  []*Embed `json:"embeds"`
}

type WebhookBuilder struct {
	url     string
	content string
	embeds  []*Embed
}

func FromUrl(url string) *WebhookBuilder {
	return &WebhookBuilder{
		url: url,
	}
}

func FromIdAndToken(id, token string) *WebhookBuilder {
	url := fmt.Sprintf("https://discord.com/api/webhooks/%s/%s", id, token)
	return FromUrl(url)
}

func (builder *WebhookBuilder) SetUrl(url string) *WebhookBuilder {
	builder.url = url
	return builder
}

func (builder *WebhookBuilder) SetContent(content string) *WebhookBuilder {
	builder.content = content
	return builder
}

func (builder *WebhookBuilder) SetEmbeds(embeds ...*Embed) *WebhookBuilder {
	builder.embeds = embeds
	return builder
}

func (builder *WebhookBuilder) Build() *Webhook {
	return &Webhook{
		url: builder.url,
		body: &Body{
			Content: builder.content,
			Embeds:  builder.embeds,
		},
	}
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
