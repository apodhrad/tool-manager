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
	toolManagerDir, err := ToolManagerMkDir(tempDir)
	assert.Nil(t, err)
	assert.Equal(t, tempDir, toolManagerDir)
}
func TestToolManagerMkDirDefault(t *testing.T) {
	tempDir := t.TempDir()
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)

	toolManagerDir, err := ToolManagerMkDir("")

	os.Setenv("HOME", oldHome)

	assert.Nil(t, err)
	assert.Equal(t, filepath.Join(tempDir, DEFAULT_TOOL_MANAGER_DIR), toolManagerDir)
	assert.DirExists(t, toolManagerDir)
	assert.Equal(t, oldHome, os.Getenv("HOME"))
}
func TestToolManagerMkDirNonexistent(t *testing.T) {
	tempDir := t.TempDir()
	os.Remove(tempDir)
	toolManagerDir, err := ToolManagerMkDir(tempDir)
	assert.Nil(t, err)
	assert.Equal(t, tempDir, toolManagerDir)
	assert.DirExists(t, tempDir)
}
func TestToolManagerMkDirError(t *testing.T) {
	tempDir := t.TempDir()
	tempFilePath := filepath.Join(tempDir, "tool-manager")
	os.Create(tempFilePath)
	toolManagerDir, err := ToolManagerMkDir(tempFilePath)
	assert.NotNil(t, err)
	assert.Empty(t, toolManagerDir)
}

func TestToolManagerLoadTools(t *testing.T) {
	err := ToolManagerLoadTools("test-resources/tool-manager")
	assert.Nil(t, err)
	tools := GetTools("", false)
	assert.Equal(t, 3, len(tools))
}
