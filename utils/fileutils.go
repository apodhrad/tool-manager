package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const DEFAULT_TOOL_MANAGER_DIR string = ".tool-manager"

func MkToolManagerDir(toolManagerDir string) (string, error) {
	if toolManagerDir == "" {
		userHomeDir, _ := os.UserHomeDir()
		toolManagerDir = filepath.Join(userHomeDir, DEFAULT_TOOL_MANAGER_DIR)
	}
	fileinfo, err := os.Stat(toolManagerDir)
	if !os.IsNotExist(err) {
		if !fileinfo.IsDir() {
			err = errors.New(fmt.Sprintf("'%s' is a file!", toolManagerDir))
		} else {
			err = nil
		}
	} else {
		err = os.Mkdir(toolManagerDir, os.ModePerm)
	}
	if err != nil {
		toolManagerDir = ""
	}
	return toolManagerDir, err
}
