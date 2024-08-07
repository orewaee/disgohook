package disgohook

type Field struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type FieldBuilder struct {
	name   string
	value  string
	inline bool
}

func NewFieldBuilder() *FieldBuilder {
	return &FieldBuilder{}
}

func (builder *FieldBuilder) Name(name string) *FieldBuilder {
	builder.name = name
	return builder
}

func (builder *FieldBuilder) Value(value string) *FieldBuilder {
	builder.value = value
	return builder
}

func (builder *FieldBuilder) Inline(inline bool) *FieldBuilder {
	builder.inline = inline
	return builder
}

func (builder *FieldBuilder) Build() *Field {
	return &Field{
		Name:   builder.name,
		Value:  builder.value,
		Inline: builder.inline,
	}
}
