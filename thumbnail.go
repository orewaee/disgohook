package disgohook

type Thumbnail struct {
	Url string `json:"url"`
}

type ThumbnailBuilder struct {
	url string
}

func NewThumbnailBuilder() *ThumbnailBuilder {
	return &ThumbnailBuilder{}
}

func (builder *ThumbnailBuilder) SetUrl(url string) *ThumbnailBuilder {
	builder.url = url
	return builder
}

func (builder *ThumbnailBuilder) Build() *Thumbnail {
	return &Thumbnail{
		Url: builder.url,
	}
}
