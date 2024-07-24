package field

type FieldBuilder struct {
	name   string
	value  string
	inline bool
}

func NewFieldBuilder() *FieldBuilder {
	return &FieldBuilder{}
}

func (builder *FieldBuilder) SetName(name string) *FieldBuilder {
	builder.name = name
	return builder
}

func (builder *FieldBuilder) SetValue(value string) *FieldBuilder {
	builder.value = value
	return builder
}

func (builder *FieldBuilder) SetInline(inline bool) *FieldBuilder {
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
