package model

type DefinitionWrapper struct {
	def Definition
}

func NewDefinitionWrapper(def Definition) *DefinitionWrapper {
	return &DefinitionWrapper{
		def: def,
	}
}

func (dw *DefinitionWrapper) UpdateDefinition(def Definition) {
	dw.def = def
}

func (dw DefinitionWrapper) Name() string {
	return dw.def.Name()
}

func (dw DefinitionWrapper) ToCSharp() string {
	return dw.def.ToCSharp()
}

func (dw DefinitionWrapper) ToVariableType() string {
	return dw.def.ToVariableType()
}

func (dw DefinitionWrapper) JsonConverter() string {
	return dw.def.JsonConverter()
}