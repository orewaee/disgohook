package disgohook

type Author struct {
	Name    string `json:"name,omitempty"`
	Url     string `json:"url,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}

type AuthorBuilder struct {
	name    string
	url     string
	iconUrl string
}

func NewAuthorBuilder() *AuthorBuilder {
	return &AuthorBuilder{}
}

func (builder *AuthorBuilder) Name(name string) *AuthorBuilder {
	builder.name = name
	return builder
}

func (builder *AuthorBuilder) Url(url string) *AuthorBuilder {
	builder.url = url
	return builder
}

func (builder *AuthorBuilder) IconUrl(iconUrl string) *AuthorBuilder {
	builder.iconUrl = iconUrl
	return builder
}

func (builder *AuthorBuilder) Build() *Author {
	return &Author{
		Name:    builder.name,
		Url:     builder.url,
		IconUrl: builder.iconUrl,
	}
}
