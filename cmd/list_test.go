package cmd

import (
	"testing"

	"github.com/apodhrad/tool-manager/utils"
	"github.com/stretchr/testify/assert"
)

const EXPECTED_OUTPUT string = `NAME     VERSION  INSTALLED  
example  1.2.1               
example  1.2.0               
example  1.1.0               
`

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
