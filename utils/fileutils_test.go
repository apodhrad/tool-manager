package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultToolManagerDir(t *testing.T) {
	assert.Equal(t, ".tool-manager", DEFAULT_TOOL_MANAGER_DIR)
}

func TestToolManagerMkDir(t *testing.T) {
	tempDir := t.TempDir()
	SetToolManagerHome(tempDir)

	toolManagerDir, err := ToolManagerMkDir()
	assert.Nil(t, err)
	assert.Equal(t, tempDir, toolManagerDir)

	UnsetToolManagerHome()
}
func TestToolManagerMkDirDefault(t *testing.T) {
	tempDir := t.TempDir()
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)

	toolManagerDir, err := ToolManagerMkDir()

	os.Setenv("HOME", oldHome)

	assert.Nil(t, err)
	assert.Equal(t, filepath.Join(tempDir, DEFAULT_TOOL_MANAGER_DIR), toolManagerDir)
	assert.DirExists(t, toolManagerDir)
	assert.Equal(t, oldHome, os.Getenv("HOME"))

	UnsetToolManagerHome()
}
func TestToolManagerMkDirNonexistent(t *testing.T) {
	tempDir := t.TempDir()
	os.Remove(tempDir)
	SetToolManagerHome(tempDir)
	toolManagerDir, err := ToolManagerMkDir()
	assert.Nil(t, err)
	assert.Equal(t, tempDir, toolManagerDir)
	assert.DirExists(t, tempDir)

	UnsetToolManagerHome()
}
func TestToolManagerMkDirError(t *testing.T) {
	tempDir := t.TempDir()
	tempFilePath := filepath.Join(tempDir, "tool-manager")
	os.Create(tempFilePath)
	SetToolManagerHome(tempFilePath)
	toolManagerDir, err := ToolManagerMkDir()
	assert.NotNil(t, err)
	assert.Empty(t, toolManagerDir)

	UnsetToolManagerHome()
}

func TestToolManagerLoadTools(t *testing.T) {
	SetToolManagerHome("test-resources/tool-manager")

	err := ToolManagerLoadTools()
	assert.Nil(t, err)
	tools := GetTools("", false)
	assert.Equal(t, 3, len(tools))

	UnsetToolManagerHome()
}
