package embed

import "github.com/orewaee/disgohook"

type Embed struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Url         string             `json:"url"`
	Color       int                `json:"color"`
	Fields      []*disgohook.Field `json:"fields"`
	Author      *Author            `json:"author"`
}
