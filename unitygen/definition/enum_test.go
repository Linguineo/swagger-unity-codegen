package definition_test

import (
	"testing"

	"github.com/recolude/swagger-unity-codegen/unitygen/definition"
	"github.com/stretchr/testify/assert"
)

func TestEnum(t *testing.T) {
	// ******************************** ARRANGE *******************************
	enum := definition.NewEnum(
		"testEnum",
		[]string{
			"A",
			"b",
			"CDF",
		},
	)

	// ********************************** ACT *********************************
	varType := enum.ToVariableType()
	name := enum.Name()
	cSharp := enum.ToCSharp()

	// ********************************* ASSERT *******************************
	assert.Equal(t, "TestEnum", varType)
	assert.Equal(t, "testEnum", name)
	assert.Equal(t, `public enum TestEnum {
	A,
	B,
	CDF
}`, cSharp)
}