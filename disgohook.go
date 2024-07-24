package disgohook

import "fmt"

func FromUrl(url string) *WebhookBuilder {
	return &WebhookBuilder{
		url: url,
	}
}

func FromIdAndToken(id, token string) *WebhookBuilder {
	url := fmt.Sprintf("https://discord.com/api/webhooks/%s/%s", id, token)
	return FromUrl(url)
}
