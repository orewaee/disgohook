package disgohook

type WebhookBuilder struct {
	url     string
	content string
}

func (builder *WebhookBuilder) SetUrl(url string) *WebhookBuilder {
	builder.url = url
	return builder
}

func (builder *WebhookBuilder) SetContent(content string) *WebhookBuilder {
	builder.content = content
	return builder
}

func (builder *WebhookBuilder) Build() *Webhook {
	return &Webhook{
		url: builder.url,
		body: &Body{
			Content: builder.content,
		},
	}
}
