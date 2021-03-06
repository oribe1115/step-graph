package search

import (
	"testing"

	"github.com/oribe1115/step-graph/lib"
	"github.com/stretchr/testify/assert"
)

var inputGraph = &lib.Graph{
	Nodes: map[int]*lib.Node{
		1: {
			ID:   1,
			Name: "a",
			Links: []*lib.Node{
				{
					ID:   2,
					Name: "b",
					Links: []*lib.Node{
						{
							ID:    3,
							Name:  "c",
							Links: []*lib.Node{},
						},
						{
							ID:    4,
							Name:  "d",
							Links: []*lib.Node{},
						},
					},
				},
			},
		},
		2: {
			ID:   2,
			Name: "b",
			Links: []*lib.Node{
				{
					ID:    3,
					Name:  "c",
					Links: []*lib.Node{},
				},
				{
					ID:    4,
					Name:  "d",
					Links: []*lib.Node{},
				},
			},
		},
		3: {
			ID:    3,
			Name:  "c",
			Links: []*lib.Node{},
		},
		4: {
			ID:    4,
			Name:  "d",
			Links: []*lib.Node{},
		},
	},
}

func TestBreadthFirstSearch(t *testing.T) {
	type input struct {
		graph      *lib.Graph
		startName  string
		targetName string
	}

	tests := []struct {
		Label          string
		Input          input
		ExpectedTraget *lib.Node
		ExpectedDepth  int
		ExpectedRoute  []*lib.Node
		IsError        bool
	}{
		{
			Label: "SUCCESS: normal",
			Input: input{
				graph:      inputGraph,
				startName:  "a",
				targetName: "d",
			},
			ExpectedTraget: &lib.Node{
				ID:    4,
				Name:  "d",
				Links: []*lib.Node{},
			},
			ExpectedDepth: 2,
			ExpectedRoute: []*lib.Node{
				{
					ID:   1,
					Name: "a",
					Links: []*lib.Node{
						{
							ID:   2,
							Name: "b",
							Links: []*lib.Node{
								{
									ID:    3,
									Name:  "c",
									Links: []*lib.Node{},
								},
								{
									ID:    4,
									Name:  "d",
									Links: []*lib.Node{},
								},
							},
						},
					},
				},
				{
					ID:   2,
					Name: "b",
					Links: []*lib.Node{
						{
							ID:    3,
							Name:  "c",
							Links: []*lib.Node{},
						},
						{
							ID:    4,
							Name:  "d",
							Links: []*lib.Node{},
						},
					},
				},
				{
					ID:    4,
					Name:  "d",
					Links: []*lib.Node{},
				},
			},
			IsError: false,
		},
		{
			Label: "SUCCESS: faild to found targetNode",
			Input: input{
				graph:      inputGraph,
				startName:  "a",
				targetName: "e",
			},
			ExpectedTraget: nil,
			ExpectedDepth:  0,
			ExpectedRoute:  nil,
			IsError:        true,
		},
		{
			Label: "FAIL: faild to found startNode",
			Input: input{
				graph:      inputGraph,
				startName:  "g",
				targetName: "d",
			},
			ExpectedTraget: nil,
			ExpectedDepth:  0,
			ExpectedRoute:  nil,
			IsError:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			target, depth, route, err := BreadthFirstSearch(test.Input.graph, test.Input.startName, test.Input.targetName)
			if test.IsError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.ExpectedTraget, target)
			assert.Equal(t, test.ExpectedDepth, depth)
			assert.Equal(t, test.ExpectedRoute, route)
		})
	}
}

func TestFindFarthermostNode(t *testing.T) {
	type input struct {
		graph     *lib.Graph
		startName string
	}
	type expected struct {
		end   *lib.Node
		depth int
		route []*lib.Node
	}

	tests := []struct {
		Label    string
		Input    input
		Expected expected
		IsError  bool
	}{
		{
			Label: "SUCCESS: normal",
			Input: input{
				graph:     inputGraph,
				startName: "a",
			},
			Expected: expected{
				end: &lib.Node{
					ID:    4,
					Name:  "d",
					Links: []*lib.Node{},
				},
				depth: 2,
				route: []*lib.Node{
					{
						ID:   1,
						Name: "a",
						Links: []*lib.Node{
							{
								ID:   2,
								Name: "b",
								Links: []*lib.Node{
									{
										ID:    3,
										Name:  "c",
										Links: []*lib.Node{},
									},
									{
										ID:    4,
										Name:  "d",
										Links: []*lib.Node{},
									},
								},
							},
						},
					},
					{
						ID:   2,
						Name: "b",
						Links: []*lib.Node{
							{
								ID:    3,
								Name:  "c",
								Links: []*lib.Node{},
							},
							{
								ID:    4,
								Name:  "d",
								Links: []*lib.Node{},
							},
						},
					},
					{
						ID:    4,
						Name:  "d",
						Links: []*lib.Node{},
					},
				},
			},
			IsError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			end, depth, route, err := FindFarthermostNode(test.Input.graph, test.Input.startName)
			if test.IsError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.Expected.end, end)
			assert.Equal(t, test.Expected.depth, depth)
			assert.Equal(t, test.Expected.route, route)
		})
	}
}
