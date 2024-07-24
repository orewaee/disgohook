package webhook

import (
	"fmt"
	"github.com/orewaee/disgohook/embed"
)

type WebhookBuilder struct {
	url     string
	content string
	embeds  []*embed.Embed
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

func (builder *WebhookBuilder) SetEmbeds(embeds ...*embed.Embed) *WebhookBuilder {
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
