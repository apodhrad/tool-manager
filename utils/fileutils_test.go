package utils

import (
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
