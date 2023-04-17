package utils

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

type Release struct {
	Version string `yaml:"version"`
	Url     string `yaml:"url"`
}

type Tool struct {
	Name       string    `yaml:"name"`
	VersionCmd string    `yaml:"version_cmd"`
	Releases   []Release `yaml:"releases"`
}

func (tool Tool) getInstalledVersion() (string, error) {
	var errOut bytes.Buffer
	cmd := exec.Command(tool.Name, strings.Split(tool.VersionCmd, " ")...)
	cmd.Stderr = &errOut
	out, err := cmd.Output()
	if err != nil {
		err = errors.New(errOut.String() + err.Error())
	}
	result := string(out)
	scanner := bufio.NewScanner(strings.NewReader(result))
	if scanner.Scan() {
		result = scanner.Text()
	}
	return result, err
}

func (tool *Tool) AddRelease(release Release) {
	tool.Releases = append(tool.Releases, release)
}

func (tool Tool) GetUrl(version string) string {
	for _, release := range tool.Releases {
		if release.Version == version {
			return release.Url
		}
	}
	return ""
}

var tools map[string]Tool

func AddTool(newTool Tool) {
	if tools == nil {
		// tools = make(map[string]map[string]string)
		tools = make(map[string]Tool)
	}
	existingTool, ok := tools[newTool.Name]
	if !ok {
		// add the whole new tool
		tools[newTool.Name] = newTool
	} else {
		// add just new releases to the existing tool
		for _, newRelease := range newTool.Releases {
			existingTool.AddRelease(newRelease)
		}
		tools[existingTool.Name] = existingTool
	}
}

func CleanTools() {
	tools = nil
}

func GetTools(name string, installed bool) map[string]Tool {
	return tools
}

func LoadToolsFromYamlFile(file string) error {
	var tools []Tool

	data, err := os.ReadFile(file)
	if err != nil {
		return errors.Join(fmt.Errorf("Cannot read yaml file '%s'", file), err)
	}

	err = yaml.Unmarshal(data, &tools)
	if err != nil {
		return errors.Join(fmt.Errorf("Cannot unmarshal yaml file '%s'", file), err)
	}

	for _, tool := range tools {
		AddTool(tool)
	}

	return err
}
