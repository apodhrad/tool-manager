package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const EXPECTED_OUTPUT string = `NAME     VERSION  INSTALLED  
example  1.2.1               
example  1.2.0               
example  1.1.0               
`

func TestListCmdRun(t *testing.T) {
	tools := getTools("", false)
	table := getTable(tools)
	output := TableToString(table)
	assert.Equal(t, EXPECTED_OUTPUT, output)
}
