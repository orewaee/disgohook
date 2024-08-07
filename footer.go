package disgohook

type Footer struct {
	Text    string `json:"text,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}

type FooterBuilder struct {
	text    string
	iconUrl string
}

func NewFooterBuilder() *FooterBuilder {
	return &FooterBuilder{}
}

func (builder *FooterBuilder) Text(text string) *FooterBuilder {
	builder.text = text
	return builder
}

func (builder *FooterBuilder) IconUrl(iconUrl string) *FooterBuilder {
	builder.iconUrl = iconUrl
	return builder
}

func (builder *FooterBuilder) Build() *Footer {
	return &Footer{
		Text:    builder.text,
		IconUrl: builder.iconUrl,
	}
}
