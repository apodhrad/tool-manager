package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const TOOL_MANAGER_HOME_ENV_KEY = "TOOL_MANAGER_HOME"
const DEFAULT_TOOL_MANAGER_DIR string = ".tool-manager"

func SetToolManagerHome(dir string) {
	os.Setenv(TOOL_MANAGER_HOME_ENV_KEY, dir)
}

func UnsetToolManagerHome() {
	os.Unsetenv(TOOL_MANAGER_HOME_ENV_KEY)
}

func ToolManagerMkDir() (string, error) {
	toolManagerHome := os.Getenv(TOOL_MANAGER_HOME_ENV_KEY)
	if toolManagerHome == "" {
		userHomeDir, _ := os.UserHomeDir()
		toolManagerHome = filepath.Join(userHomeDir, DEFAULT_TOOL_MANAGER_DIR)
	}
	fileinfo, err := os.Stat(toolManagerHome)
	if !os.IsNotExist(err) {
		if !fileinfo.IsDir() {
			err = errors.New(fmt.Sprintf("'%s' is a file!", toolManagerHome))
		} else {
			err = nil
		}
	} else {
		err = os.Mkdir(toolManagerHome, os.ModePerm)
	}
	if err != nil {
		toolManagerHome = ""
	}
	return toolManagerHome, err
}

func ToolManagerLoadTools() error {
	CleanTools()
	toolManagerDir, err := ToolManagerMkDir()
	if err == nil {
		err = filepath.Walk(toolManagerDir, func(path string, info fs.FileInfo, err error) error {
			if err == nil && strings.HasSuffix(path, ".yaml") {
				err = LoadToolsFromYamlFile(path)
			}
			return err
		})
	}
	return err
}
