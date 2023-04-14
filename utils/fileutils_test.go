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

func TestMkToolManagerDir(t *testing.T) {
	tempDir := t.TempDir()
	toolManagerDir, err := MkToolManagerDir(tempDir)
	assert.Nil(t, err)
	assert.Equal(t, tempDir, toolManagerDir)
}
func TestMkToolManagerDirDefault(t *testing.T) {
	tempDir := t.TempDir()
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)

	toolManagerDir, err := MkToolManagerDir("")

	os.Setenv("HOME", oldHome)

	assert.Nil(t, err)
	assert.Equal(t, filepath.Join(tempDir, DEFAULT_TOOL_MANAGER_DIR), toolManagerDir)
	assert.DirExists(t, toolManagerDir)
	assert.Equal(t, oldHome, os.Getenv("HOME"))
}
func TestMkToolManagerDirNonexistent(t *testing.T) {
	tempDir := t.TempDir()
	os.Remove(tempDir)
	toolManagerDir, err := MkToolManagerDir(tempDir)
	assert.Nil(t, err)
	assert.Equal(t, tempDir, toolManagerDir)
	assert.DirExists(t, tempDir)
}
func TestMkToolManagerDirError(t *testing.T) {
	tempDir := t.TempDir()
	tempFilePath := filepath.Join(tempDir, "tool-manager")
	os.Create(tempFilePath)
	toolManagerDir, err := MkToolManagerDir(tempFilePath)
	assert.NotNil(t, err)
	assert.Empty(t, toolManagerDir)
}
