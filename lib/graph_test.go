package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitGraph(t *testing.T) {
	tests := []struct {
		Label    string
		Expected *Graph
	}{
		{
			Label:    "SUCCESS: normal",
			Expected: &Graph{},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := InitGraph()
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestGraphFindNode(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *Graph
		Input    int
		Expected *Node
	}{
		{
			Label: "SUCCESS: normal",
			Use: &Graph{
				Nodes: map[int]*Node{
					1: &Node{
						ID:    1,
						Name:  "a",
						Links: []*Node{},
					},
					2: &Node{
						ID:    2,
						Name:  "b",
						Links: []*Node{},
					},
				},
			},
			Input: 1,
			Expected: &Node{
				ID:    1,
				Name:  "a",
				Links: []*Node{},
			},
		},
		{
			Label: "FAIL: not found",
			Use: &Graph{
				Nodes: map[int]*Node{
					1: &Node{
						ID:    1,
						Name:  "a",
						Links: []*Node{},
					},
					2: &Node{
						ID:    2,
						Name:  "b",
						Links: []*Node{},
					},
				},
			},
			Expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := test.Use.FindNode(test.Input)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestGraphSetNode(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *Graph
		Input    *Node
		Expected *Graph
		IsError  bool
	}{
		{
			Label: "SUCCESS: normal",
			Use: &Graph{
				Nodes: map[int]*Node{
					1: {
						ID:    1,
						Name:  "a",
						Links: []*Node{},
					},
					2: {
						ID:    2,
						Name:  "b",
						Links: []*Node{},
					},
				},
			},
			Input: &Node{
				ID:    3,
				Name:  "c",
				Links: []*Node{},
			},
			Expected: &Graph{
				Nodes: map[int]*Node{
					1: {
						ID:    1,
						Name:  "a",
						Links: []*Node{},
					},
					2: {
						ID:    2,
						Name:  "b",
						Links: []*Node{},
					},
					3: {
						ID:    3,
						Name:  "c",
						Links: []*Node{},
					},
				},
			},
			IsError: false,
		},
		{
			Label: "FAIL: deplicate id",
			Use: &Graph{
				Nodes: map[int]*Node{
					1: {
						ID:    1,
						Name:  "a",
						Links: []*Node{},
					},
					2: {
						ID:    2,
						Name:  "b",
						Links: []*Node{},
					},
				},
			},
			Input: &Node{
				ID:    1,
				Name:  "c",
				Links: []*Node{},
			},
			Expected: nil,
			IsError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			err := test.Use.SetNode(test.Input)
			if test.IsError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, test.Expected, test.Use)
		})
	}
}

func TestGraphLinkNodes(t *testing.T) {
	type input struct {
		fromID int
		toID   int
	}

	tests := []struct {
		Label    string
		Use      *Graph
		Input    input
		Expected *Graph
		IsError  bool
	}{
		{
			Label: "SUCCESS: normal",
			Use: &Graph{
				Nodes: map[int]*Node{
					1: {
						ID:    1,
						Name:  "a",
						Links: []*Node{},
					},
					2: {
						ID:    2,
						Name:  "b",
						Links: []*Node{},
					},
				},
			},
			Input: input{
				fromID: 1,
				toID:   2,
			},
			Expected: &Graph{
				Nodes: map[int]*Node{
					1: {
						ID:   1,
						Name: "a",
						Links: []*Node{
							{
								ID:    2,
								Name:  "b",
								Links: []*Node{},
							},
						},
					},
					2: {
						ID:    2,
						Name:  "b",
						Links: []*Node{},
					},
				},
			},
			IsError: false,
		},
		{
			Label: "FAIL: not found fromNode",
			Use: &Graph{
				Nodes: map[int]*Node{
					1: {
						ID:    1,
						Name:  "a",
						Links: []*Node{},
					},
					2: {
						ID:    2,
						Name:  "b",
						Links: []*Node{},
					},
				},
			},
			Input: input{
				fromID: 3,
				toID:   2,
			},
			Expected: nil,
			IsError:  true,
		},
		{
			Label: "FAIL: not found toNode",
			Use: &Graph{
				Nodes: map[int]*Node{
					1: {
						ID:    1,
						Name:  "a",
						Links: []*Node{},
					},
					2: {
						ID:    2,
						Name:  "b",
						Links: []*Node{},
					},
				},
			},
			Input: input{
				fromID: 1,
				toID:   3,
			},
			Expected: nil,
			IsError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			err := test.Use.LinkNodes(test.Input.fromID, test.Input.toID)
			if test.IsError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, test.Expected, test.Use)
		})
	}
}
