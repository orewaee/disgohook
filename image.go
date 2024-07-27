package disgohook

type Image struct {
	Url string `json:"url"`
}

type ImageBuilder struct {
	url string
}

func NewImageBuilder() *ImageBuilder {
	return &ImageBuilder{}
}

func (builder *ImageBuilder) SetUrl(url string) *ImageBuilder {
	builder.url = url
	return builder
}

func (builder *ImageBuilder) Build() *Image {
	return &Image{
		Url: builder.url,
	}
}
