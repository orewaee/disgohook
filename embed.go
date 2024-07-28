package disgohook

import "time"

type Embed struct {
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Url         string     `json:"url,omitempty"`
	Color       int        `json:"color,omitempty"`
	Image       *Image     `json:"image,omitempty"`
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`
	Fields      []*Field   `json:"fields,omitempty"`
	Author      *Author    `json:"author,omitempty"`
	Footer      *Footer    `json:"footer,omitempty"`
	Timestamp   *time.Time `json:"timestamp,omitempty"`
}

type EmbedBuilder struct {
	title       string
	description string
	url         string
	color       int
	image       *Image
	thumbnail   *Thumbnail
	fields      []*Field
	author      *Author
	footer      *Footer
	timestamp   *time.Time
}

func NewEmbedBuilder() *EmbedBuilder {
	return &EmbedBuilder{}
}

func (builder *EmbedBuilder) SetTitle(title string) *EmbedBuilder {
	builder.title = title
	return builder
}

func (builder *EmbedBuilder) SetDescription(description string) *EmbedBuilder {
	builder.description = description
	return builder
}

func (builder *EmbedBuilder) SetUrl(url string) *EmbedBuilder {
	builder.url = url
	return builder
}

func (builder *EmbedBuilder) SetColor(color int) *EmbedBuilder {
	builder.color = color
	return builder
}

func (builder *EmbedBuilder) SetImage(image *Image) *EmbedBuilder {
	builder.image = image
	return builder
}

func (builder *EmbedBuilder) SetThumbnail(thumbnail *Thumbnail) *EmbedBuilder {
	builder.thumbnail = thumbnail
	return builder
}

func (builder *EmbedBuilder) SetFields(fields ...*Field) *EmbedBuilder {
	builder.fields = fields
	return builder
}

func (builder *EmbedBuilder) SetAuthor(author *Author) *EmbedBuilder {
	builder.author = author
	return builder
}

func (builder *EmbedBuilder) SetFooter(footer *Footer) *EmbedBuilder {
	builder.footer = footer
	return builder
}

func (builder *EmbedBuilder) SetTimestamp(timestamp *time.Time) *EmbedBuilder {
	builder.timestamp = timestamp
	return builder
}

func (builder *EmbedBuilder) Build() *Embed {
	return &Embed{
		Title:       builder.title,
		Description: builder.description,
		Url:         builder.url,
		Color:       builder.color,
		Fields:      builder.fields,
		Author:      builder.author,
	}
}
