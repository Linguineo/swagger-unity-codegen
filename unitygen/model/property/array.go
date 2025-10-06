package property

import (
	"fmt"
	"strings"

	"github.com/recolude/swagger-unity-codegen/unitygen/convention"
	"github.com/recolude/swagger-unity-codegen/unitygen/model"
)

type Array struct {
	name string
	prop model.Property
}

func NewArray(name string, prop model.Property) Array {
	return Array{
		name: name,
		prop: prop,
	}
}

func (sp Array) Name() string {
	return sp.name
}

func (sp Array) Property() model.Property {
	return sp.prop
}

func (sp Array) ToVariableType() string {
	return fmt.Sprintf("List<%s>", sp.prop.ToVariableType())
}

func (sp Array) EmptyValue() string {
	return "null"
}

func (sp Array) ClassVariables(className string) string {
    propNameToUse := convention.TitleCase(sp.name)
    
    // Check for conflict: If the generated property name equals the parent class name.
    if propNameToUse == className {
        // Fix: Prepend an underscore to the property name (e.g., UserInterests -> _UserInterests).
        propNameToUse = "_" + propNameToUse
    }
	
	builder := strings.Builder{}
	builder.WriteString("\t[JsonProperty(\"")
	builder.WriteString(sp.name)
	builder.WriteString("\")]\n\tpublic ")
	builder.WriteString(sp.ToVariableType())
	builder.WriteString(" ")
	builder.WriteString(propNameToUse)
	builder.WriteString(" { get; set; }\n")
	return builder.String()
}
