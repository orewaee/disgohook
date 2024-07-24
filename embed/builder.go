package embed

import (
	"github.com/orewaee/disgohook"
	"github.com/orewaee/disgohook/author"
)

type EmbedBuilder struct {
	title       string
	description string
	url         string
	color       int
	fields      []*disgohook.Field
	author      *author.Author
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

func (builder *EmbedBuilder) SetFields(fields ...*disgohook.Field) *EmbedBuilder {
	builder.fields = fields
	return builder
}

func (builder *EmbedBuilder) SetAuthor(author *author.Author) *EmbedBuilder {
	builder.author = author
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
