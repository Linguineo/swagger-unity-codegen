package property

import (
	"fmt"
	"strings"

	"github.com/recolude/swagger-unity-codegen/unitygen/convention"
)

type String struct {
	name   string
	format string
}

func NewString(name string, format string) String {
	return String{
		name:   name,
		format: format,
	}
}

func (sp String) Name() string {
	return sp.name
}

func (sp String) ToVariableType() string {
	switch sp.format {
	case "date-time":
		return "System.DateTime"

	default:
		return "string"
	}
}

func (sp String) EmptyValue() string {
	return "null"
}

func (sp String) ClassVariables(className string) string {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "\t[JsonProperty(\"%s\")]\n", sp.Name())

	propNameToUse := convention.TitleCase(sp.Name())
    if propNameToUse == className {
        propNameToUse = propNameToUse + "_";
    }

	switch sp.format {
	case "date-time":
		fmt.Fprintf(&builder, "\tpublic string %s;\n\n", convention.CamelCase(sp.Name()))
		fmt.Fprintf(&builder, "\tpublic System.DateTime? %s { get { if (%s != null) { return System.DateTime.Parse(%s); } else { return null; } } }\n", propNameToUse, convention.CamelCase(sp.Name()), convention.CamelCase(sp.Name()))
		break

	default:
		fmt.Fprintf(&builder, "\tpublic string %s { get; set; }\n", propNameToUse)
		break
	}

	return builder.String()
}
