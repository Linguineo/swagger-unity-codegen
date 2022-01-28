package property

import (
	"strings"

	"github.com/recolude/swagger-unity-codegen/unitygen/convention"
	"github.com/recolude/swagger-unity-codegen/unitygen/model"
)

type Object struct {
	name string
	obj  model.Object
}

func NewObject(name string, obj model.Object) Object {
	return Object{
		name: name,
		obj:  obj,
	}
}

// Object is the actual object definition of the property
func (op Object) Object() model.Object {
	return op.obj
}

// Name of the property (generally a c# variable name)
func (op Object) Name() string {
	return op.name
}

// ToVariableType returns the name of the variable type that exists in c# (ie: float, int,s tring)
func (op Object) ToVariableType() string {
	return op.obj.ToVariableType()
}

// EmptyValue is the value that represents the property has yet to be set.
func (op Object) EmptyValue() string {
	return "null"
}

// What gets written to the c# class definition.
func (op Object) ClassVariables(className string) string {
	builder := strings.Builder{}
	builder.WriteString("\t")
	builder.WriteString(op.obj.ToCSharp())
	builder.WriteString("\n\t[JsonProperty(\"")
	builder.WriteString(op.name)
	builder.WriteString("\")]\n\tpublic ")
	builder.WriteString(op.ToVariableType())
	builder.WriteString(" ")
	builder.WriteString(convention.TitleCase(op.name))
	builder.WriteString(" { get; set; }\n")
	return builder.String()
}
