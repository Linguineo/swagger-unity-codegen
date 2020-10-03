package main

import (
	"os"
	"strings"
	"testing"

	"github.com/recolude/swagger-unity-codegen/unitygen"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestFilterServiceByTags_DoesNothingWithNoTags(t *testing.T) {
	// ******************************** ARRANGE *******************************
	spec := unitygen.NewSpec(unitygen.SpecInfo{}, nil, nil, []unitygen.Service{
		unitygen.NewService("A", nil),
		unitygen.NewService("B", nil),
	})

	// ********************************** ACT *********************************
	filterSpecForTags(spec, nil)

	// ********************************* ASSERT *******************************
	if assert.Len(t, spec.Services, 2) {
		assert.Equal(t, spec.Services[0].Name(), "A")
		assert.Equal(t, spec.Services[1].Name(), "B")
	}
}

func TestFilterServiceByTags_Filters(t *testing.T) {
	// ******************************** ARRANGE *******************************
	spec := unitygen.NewSpec(unitygen.SpecInfo{}, nil, nil, []unitygen.Service{
		unitygen.NewService("A", nil),
		unitygen.NewService("B", nil),
	})

	// ********************************** ACT *********************************
	filterSpecForTags(spec, []string{"A"})

	// ********************************* ASSERT *******************************
	if assert.Len(t, spec.Services, 2) {
		assert.Equal(t, spec.Services[0].Name(), "A")
	}
}

func TestNoNamespace(t *testing.T) {
	// ******************************** ARRANGE *******************************
	appFS := afero.NewMemMapFs()
	afero.WriteFile(appFS, "swagger.json", []byte("{ }"), os.ModePerm)

	out := strings.Builder{}
	errOut := strings.Builder{}
	app := buildApp(appFS, &out, &errOut)

	// ********************************** ACT *********************************
	err := app.Run([]string{"swag3d", "--file", "swagger.json", "generate"})

	// ********************************* ASSERT *******************************
	assert.NoError(t, err)
	assert.Equal(t, "", errOut.String())
	assert.Equal(t, `// This code was generated by: 
// https://github.com/recolude/swagger-unity-codegen
// Issues and PRs welcome :)

using UnityEngine;
using UnityEngine.Networking;
using System.Collections;

#region Definitions

#endregion

#region Services

public interface Config {

	// The base URL to which the endpoint paths are appended
	string BasePath { get; set; }

}

[System.Serializable]
[CreateAssetMenu(menuName = "", fileName = "")]
public class : ScriptableObject, Config {

	// The base URL to which the endpoint paths are appended
	[SerializeField]
	public string BasePath { get; set; }

	public (string basePath) {
		this.BasePath = basePath;
	}

}

#endregion

`, out.String())
}

func TestWithNamespace(t *testing.T) {
	// ******************************** ARRANGE *******************************
	appFS := afero.NewMemMapFs()
	afero.WriteFile(appFS, "swagger.json", []byte("{ }"), os.ModePerm)

	out := strings.Builder{}
	errOut := strings.Builder{}
	app := buildApp(appFS, &out, &errOut)

	// ********************************** ACT *********************************
	err := app.Run([]string{"swag3d", "--file", "swagger.json", "generate", "--namespace", "example"})

	// ********************************* ASSERT *******************************
	assert.NoError(t, err)
	assert.Equal(t, "", errOut.String())
	assert.Equal(t, `// This code was generated by: 
// https://github.com/recolude/swagger-unity-codegen
// Issues and PRs welcome :)

using UnityEngine;
using UnityEngine.Networking;
using System.Collections;

namespace Example {

#region Definitions

#endregion

#region Services

public interface Config {

	// The base URL to which the endpoint paths are appended
	string BasePath { get; set; }

}

[System.Serializable]
[CreateAssetMenu(menuName = "", fileName = "")]
public class : ScriptableObject, Config {

	// The base URL to which the endpoint paths are appended
	[SerializeField]
	public string BasePath { get; set; }

	public (string basePath) {
		this.BasePath = basePath;
	}

}

#endregion

}`, out.String())
}