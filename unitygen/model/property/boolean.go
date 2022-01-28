package property

import (
	"fmt"

	"github.com/recolude/swagger-unity-codegen/unitygen/convention"
)

type Boolean struct {
	name string
}

func NewBoolean(name string) Boolean {
	return Boolean{
		name: name,
	}
}

func (sp Boolean) Name() string {
	return sp.name
}

func (sp Boolean) ToVariableType() string {
	return "bool"
}

func (sp Boolean) EmptyValue() string {
	return "false"
}

func (sp Boolean) ClassVariables(className string) string {
    propNameToUse := convention.TitleCase(sp.Name())
    if propNameToUse == className {
        propNameToUse = propNameToUse + "_";
    }

	return fmt.Sprintf("\t[JsonProperty(\"%s\")]\n\tpublic %s %s { get; private set; }\n", sp.Name(), sp.ToVariableType(), propNameToUse)
}
