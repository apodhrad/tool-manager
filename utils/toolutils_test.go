package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() {
	CleanTools()
}

func teardown() {
	CleanTools()
}

func TestEmptyTools(t *testing.T) {
	setup()

	tools := GetTools("", false)
	assert.Empty(t, tools)

	teardown()
}
func TestAddingNewTool(t *testing.T) {
	setup()

	example1 := Tool{Name: "example1"}
	example2 := Tool{Name: "example2"}

	AddTool(example1)
	AddTool(example2)

	tools := GetTools("", false)
	assert.Equal(t, 2, len(tools))
	assert.Equal(t, example1, tools["example1"])
	assert.Equal(t, example2, tools["example2"])

	example3 := Tool{Name: "example2"}

	AddTool(example3)
	tools = GetTools("", false)
	assert.Equal(t, 2, len(tools))
	assert.Equal(t, example1, tools["example1"])
	assert.Equal(t, example3, tools["example2"])

	releases := []Release{{Version: "1.0.0"}, {Version: "1.1.0"}}
	example4 := Tool{Name: "example2", Releases: releases}

	AddTool(example4)
	tools = GetTools("", false)
	assert.Equal(t, 2, len(tools))
	assert.Equal(t, example1, tools["example1"])
	assert.Equal(t, example4, tools["example2"])

	teardown()
}

func TestGettingUrl(t *testing.T) {
	releases := []Release{
		{Version: "1.0.0", Url: "example.com/release-1.0.0"},
		{Version: "1.1.0", Url: "example.com/release-1.1.0"},
	}
	example := Tool{Name: "example", Releases: releases}

	assert.Equal(t, 2, len(example.Releases))
	assert.Equal(t, "example.com/release-1.0.0", example.GetUrl("1.0.0"))
	assert.Equal(t, "example.com/release-1.1.0", example.GetUrl("1.1.0"))
	assert.Equal(t, "", example.GetUrl("1.1.1"))
}

func TestGettingInstalledVersion(t *testing.T) {
	example := Tool{Name: "jq", VersionCmd: "--version"}
	version, err := example.getInstalledVersion()

	assert.Nil(t, err)
	assert.Equal(t, "jq-1.6", version)
}

func TestGettingInstalledVersionError(t *testing.T) {
	example := Tool{Name: "foo", VersionCmd: "--version"}
	version, err := example.getInstalledVersion()

	assert.NotNil(t, err)
	assert.Empty(t, version)
	assert.Equal(t, "exec: \"foo\": executable file not found in $PATH", err.Error())
}
