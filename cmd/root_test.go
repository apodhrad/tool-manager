package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() {
}

func teardown() {
}

func TestEmptyTable(t *testing.T) {
	expected := `HEADERA  HEADERB  HEADERC  
`
	tbl := NewTable("HeaderA", "HeaderB", "HeaderC")
	actual := TableToString(tbl)
	assert.Equal(t, expected, actual)
}

func TestNonEmptyTable(t *testing.T) {
	expected := `HEADERA          HEADERB  HEADERC  
ValueA1          ValueB1  ValueC1  
VeryLongValueA2  ValueB2  ValueC2  
`
	tbl := NewTable("HeaderA", "HeaderB", "HeaderC")
	tbl.AddRow("ValueA1", "ValueB1", "ValueC1")
	tbl.AddRow("VeryLongValueA2", "ValueB2", "ValueC2")
	actual := TableToString(tbl)
	assert.Equal(t, expected, actual)
}
