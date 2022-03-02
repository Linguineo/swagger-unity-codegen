package property

import (
	"strings"

	"github.com/recolude/swagger-unity-codegen/unitygen/convention"
)

type Integer struct {
	name   string
	format string
}

func NewInteger(name string, format string) Integer {
	return Integer{
		name:   name,
		format: format,
	}
}

func (sp Integer) Name() string {
	return sp.name
}

func (sp Integer) ToVariableType() string {
	if sp.format == "int32" {
        return "int?"
    }

    if sp.format == "int64" {
        return "long?"
    }

    return "int?"
}

func (sp Integer) EmptyValue() string {
	switch sp.format {
	default:
		return "0"
	}
}

func (sp Integer) ClassVariables(className string) string {
	builder := strings.Builder{}
	builder.WriteString("	[JsonProperty(\"")
	builder.WriteString(sp.Name())
	builder.WriteString("\")]\n\tpublic ")
	builder.WriteString(sp.ToVariableType())
	builder.WriteString(" ")
	builder.WriteString(convention.TitleCase(sp.Name()))
	builder.WriteString(" { get; set; }\n")
	return builder.String()
}
