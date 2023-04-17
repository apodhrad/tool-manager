package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const DEFAULT_TOOL_MANAGER_DIR string = ".tool-manager"

func ToolManagerMkDir(toolManagerDir string) (string, error) {
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

func ToolManagerLoadTools(dir string) error {
	CleanTools()
	toolManagerDir, err := ToolManagerMkDir(dir)
	if err == nil {
		filepath.Walk(toolManagerDir, func(path string, info fs.FileInfo, err error) error {
			if err == nil && strings.HasSuffix(path, ".yaml") {
				err = LoadToolsFromYamlFile(path)
			}
			return err
		})
	}
	return err
}
