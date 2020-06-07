package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadNodeData(t *testing.T) {
	tests := []struct {
		Label    string
		Input    string
		Expected []NodeData
		IsError  bool
	}{
		{
			Label: "SUCCESS: normal",
			Input: "./testdata/nodes.txt",
			Expected: []NodeData{
				{0, "a"},
				{1, "b"},
				{2, "c"},
			},
			IsError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got, err := ReadNodeData(test.Input)
			if test.IsError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestReadLinkData(t *testing.T) {
	tests := []struct {
		Label    string
		Input    string
		Expected []LinkData
		IsError  bool
	}{
		{
			Label: "SUCCESS: normal",
			Input: "./testdata/links.txt",
			Expected: []LinkData{
				{0, 1}, {0, 2}, {1, 2},
			},
			IsError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got, err := ReadLinkData(test.Input)
			if test.IsError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.Expected, got)
		})
	}
}
