package author

type AuthorBuilder struct {
	name    string
	url     string
	iconUrl string
}

func NewAuthorBuilder() *AuthorBuilder {
	return &AuthorBuilder{}
}

func (builder *AuthorBuilder) SetName(name string) *AuthorBuilder {
	builder.name = name
	return builder
}

func (builder *AuthorBuilder) SetUrl(url string) *AuthorBuilder {
	builder.url = url
	return builder
}

func (builder *AuthorBuilder) SetIconUrl(iconUrl string) *AuthorBuilder {
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
