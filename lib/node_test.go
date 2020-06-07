package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNode(t *testing.T) {
	type input struct {
		ID   int
		Name string
	}
	tests := []struct {
		Label    string
		Input    input
		Expected *Node
		Error    bool
	}{
		{Label: "SUCCESS: normal",
			Input: input{
				ID:   1,
				Name: "a",
			},
			Expected: &Node{
				ID:   1,
				Name: "a",
			},
			Error: false},
		{Label: "FAIL: name is blank",
			Input:    input{ID: 2, Name: ""},
			Expected: nil,
			Error:    true},
	}

	for _, test := range tests {
		got, err := CreateNode(test.Input.ID, test.Input.Name)
		if test.Error {
			assert.Error(t, err, test.Label)
			return
		}

		assert.NoError(t, err, test.Label)
		assert.Equal(t, test.Expected, got, test.Label)
	}
}
