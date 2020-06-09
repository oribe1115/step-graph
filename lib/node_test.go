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
		IsError  bool
	}{
		{
			Label: "SUCCESS: normal",
			Input: input{
				ID:   1,
				Name: "a",
			},
			Expected: &Node{
				ID:   1,
				Name: "a",
			},
			IsError: false,
		},
		{
			Label:    "FAIL: name is blank",
			Input:    input{ID: 2, Name: ""},
			Expected: nil,
			IsError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got, err := CreateNode(test.Input.ID, test.Input.Name)
			if test.IsError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestNodeLinkTo(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *Node
		Input    *Node
		Expected *Node
	}{
		{
			Label: "SUCCESS: normal",
			Use: &Node{
				ID:   1,
				Name: "a",
				Links: []*Node{
					{
						2,
						"b",
						[]*Node{},
					},
					{
						3,
						"c",
						[]*Node{
							{
								2,
								"b",
								[]*Node{},
							},
						},
					},
				},
			},
			Input: &Node{
				4,
				"d",
				[]*Node{},
			},
			Expected: &Node{
				ID:   1,
				Name: "a",
				Links: []*Node{
					{
						2,
						"b",
						[]*Node{},
					},
					{
						3,
						"c",
						[]*Node{
							{
								2,
								"b",
								[]*Node{},
							},
						},
					},
					{
						4,
						"d",
						[]*Node{},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			test.Use.LinkTo(test.Input)
			assert.Equal(t, test.Expected, test.Use)
		})
	}
}

func TestNodeSprint(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *Node
		Expected string
	}{
		{
			Label: "SUCCESS: normal",
			Use: &Node{
				ID:    1,
				Name:  "a",
				Links: []*Node{},
			},
			Expected: "{1: a}",
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := test.Use.Sprint()
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestSprintNodeListAsRoute(t *testing.T) {
	tests := []struct {
		Label    string
		Input    []*Node
		Expected string
	}{
		{
			Label: "SUCCESS: normal",
			Input: []*Node{
				{
					ID:    1,
					Name:  "a",
					Links: []*Node{},
				},
				{
					ID:    2,
					Name:  "b",
					Links: []*Node{},
				},
			},
			Expected: "{1: a} -> {2: b}",
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := SprintNodeListAsRoute(test.Input)
			assert.Equal(t, test.Expected, got)
		})
	}
}
