package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/apodhrad/tool-manager/utils"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

const EXPECTED_OUTPUT string = `NAME     VERSION  INSTALLED  
example  1.2.1               
example  1.2.0               
example  1.1.0               
`

func TestListCmdRunE(t *testing.T) {
	tempDir := t.TempDir()
	utils.SetToolManagerHome(tempDir)

	listCmdRunE := listCmd.RunE
	assert.NotNil(t, listCmdRunE)

	err := listCmdRunE(&cobra.Command{}, []string{})
	assert.Nil(t, err)
}

func TestListCmdRunEWithError(t *testing.T) {
	tempDir := t.TempDir()
	data := "name: foo"
	os.WriteFile(filepath.Join(tempDir, "test.yaml"), []byte(data), os.ModePerm)
	utils.SetToolManagerHome(tempDir)

	listCmdRunE := listCmd.RunE
	assert.NotNil(t, listCmdRunE)

	err := listCmdRunE(&cobra.Command{}, []string{})
	assert.NotNil(t, err)
}

func TestToolsToString(t *testing.T) {
	example := utils.Tool{
		Name: "example", Releases: []utils.Release{
			{Version: "1.1.0"},
			{Version: "1.2.0"},
			{Version: "1.2.1"},
		},
	}
	tools := map[string]utils.Tool{"example": example}
	output := ToolsToString(tools)

	assert.NotNil(t, EXPECTED_OUTPUT, output)
}
