package search

import (
	"testing"

	"github.com/oribe1115/step-graph/lib"
	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	type input struct {
		graph      *lib.Graph
		startName  string
		targetName string
		getCost    func(from *lib.Node, to *lib.Node) (int, error)
	}
	type expected struct {
		target    *lib.Node
		totalCost int
		route     []*lib.Node
	}
	tests := []struct {
		Label    string
		Input    input
		Expected expected
		IsError  bool
	}{
		{
			Label: "SUCCSESS: normal",
			Input: input{
				graph: &lib.Graph{
					Nodes: map[int]*lib.Node{
						0: {
							ID:   0,
							Name: "a",
							Links: []*lib.Node{
								{
									ID:   1,
									Name: "b",
									Links: []*lib.Node{
										{
											ID:   3,
											Name: "d",
											Links: []*lib.Node{
												{
													ID:    4,
													Name:  "e",
													Links: []*lib.Node{},
												},
											},
										},
									},
								},
								{
									ID:   2,
									Name: "c",
									Links: []*lib.Node{
										{
											ID:    4,
											Name:  "e",
											Links: []*lib.Node{},
										},
									},
								},
								{
									ID:   3,
									Name: "d",
									Links: []*lib.Node{
										{
											ID:    4,
											Name:  "e",
											Links: []*lib.Node{},
										},
									},
								},
							},
						},
						1: {
							ID:   1,
							Name: "b",
							Links: []*lib.Node{
								{
									ID:   3,
									Name: "d",
									Links: []*lib.Node{
										{
											ID:    4,
											Name:  "e",
											Links: []*lib.Node{},
										},
									},
								},
							},
						},
						2: {
							ID:   2,
							Name: "c",
							Links: []*lib.Node{
								{
									ID:    4,
									Name:  "e",
									Links: []*lib.Node{},
								},
							},
						},
						3: {
							ID:   3,
							Name: "d",
							Links: []*lib.Node{
								{
									ID:    4,
									Name:  "e",
									Links: []*lib.Node{},
								},
							},
						},
						4: {
							ID:    4,
							Name:  "e",
							Links: []*lib.Node{},
						},
					},
				},
				startName:  "a",
				targetName: "e",
				getCost: func(from *lib.Node, to *lib.Node) (int, error) {
					return from.ID + to.ID, nil
				},
			},
			Expected: expected{
				target: &lib.Node{
					ID:    4,
					Name:  "e",
					Links: []*lib.Node{},
				},
				totalCost: 8,
				route: []*lib.Node{
					{
						ID:   0,
						Name: "a",
						Links: []*lib.Node{
							{
								ID:   1,
								Name: "b",
								Links: []*lib.Node{
									{
										ID:   3,
										Name: "d",
										Links: []*lib.Node{
											{
												ID:    4,
												Name:  "e",
												Links: []*lib.Node{},
											},
										},
									},
								},
							},
							{
								ID:   2,
								Name: "c",
								Links: []*lib.Node{
									{
										ID:    4,
										Name:  "e",
										Links: []*lib.Node{},
									},
								},
							},
							{
								ID:   3,
								Name: "d",
								Links: []*lib.Node{
									{
										ID:    4,
										Name:  "e",
										Links: []*lib.Node{},
									},
								},
							},
						},
					},
					{
						ID:   2,
						Name: "c",
						Links: []*lib.Node{
							{
								ID:    4,
								Name:  "e",
								Links: []*lib.Node{},
							},
						},
					},
					{
						ID:    4,
						Name:  "e",
						Links: []*lib.Node{},
					},
				},
			},
			IsError: false,
		},
		{
			Label: "FAIL: firstNode not found",
			Input: input{
				graph: &lib.Graph{
					Nodes: map[int]*lib.Node{
						0: {ID: 0, Name: "a", Links: []*lib.Node{}},
					},
				},
				startName:  "b",
				targetName: "a",
				getCost:    nil,
			},
			Expected: expected{
				target:    nil,
				totalCost: 0,
				route:     nil,
			},
			IsError: true,
		},
		{
			Label: "FAIL: targetNode not found",
			Input: input{
				graph: &lib.Graph{
					Nodes: map[int]*lib.Node{
						0: {ID: 0, Name: "a", Links: []*lib.Node{}},
					},
				},
				startName:  "a",
				targetName: "b",
				getCost:    nil,
			},
			Expected: expected{
				target:    nil,
				totalCost: 0,
				route:     nil,
			},
			IsError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			target, totalCost, route, err := Dijkstra(test.Input.graph, test.Input.startName, test.Input.targetName, test.Input.getCost)
			if test.IsError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err, err)
			assert.Equal(t, test.Expected.target, target)
			assert.Equal(t, test.Expected.totalCost, totalCost)
			assert.Equal(t, test.Expected.route, route)
		})
	}
}

func TestJustCostRouteWithDijkstra(t *testing.T) {
	type input struct {
		graph      *lib.Graph
		startName  string
		targetCost int
		getCost    func(from *lib.Node, to *lib.Node) (int, error)
	}
	tests := []struct {
		Label    string
		Input    input
		Expected [][]*lib.Node
		IsError  bool
	}{
		{
			Label: "SUCCESS: normal",
			Input: input{
				graph: &lib.Graph{
					Nodes: map[int]*lib.Node{
						0: {
							ID:   0,
							Name: "a",
							Links: []*lib.Node{
								{
									ID:   1,
									Name: "b",
									Links: []*lib.Node{
										{
											ID:    3,
											Name:  "c",
											Links: []*lib.Node{},
										},
									},
								},
								{
									ID:    5,
									Name:  "d",
									Links: []*lib.Node{},
								},
							},
						},
						1: {
							ID:   1,
							Name: "b",
							Links: []*lib.Node{
								{
									ID:    3,
									Name:  "c",
									Links: []*lib.Node{},
								},
							},
						},
						3: {
							ID:    3,
							Name:  "c",
							Links: []*lib.Node{},
						},
						5: {
							ID:    5,
							Name:  "d",
							Links: []*lib.Node{},
						},
					},
				},
				startName:  "a",
				targetCost: 5,
				getCost: func(from *lib.Node, to *lib.Node) (int, error) {
					return from.ID + to.ID, nil
				},
			},
			Expected: [][]*lib.Node{
				{
					{
						ID:   0,
						Name: "a",
						Links: []*lib.Node{
							{
								ID:   1,
								Name: "b",
								Links: []*lib.Node{
									{
										ID:    3,
										Name:  "c",
										Links: []*lib.Node{},
									},
								},
							},
							{
								ID:    5,
								Name:  "d",
								Links: []*lib.Node{},
							},
						},
					},
					{
						ID:    5,
						Name:  "d",
						Links: []*lib.Node{},
					},
				},
				{
					{
						ID:   0,
						Name: "a",
						Links: []*lib.Node{
							{
								ID:   1,
								Name: "b",
								Links: []*lib.Node{
									{
										ID:    3,
										Name:  "c",
										Links: []*lib.Node{},
									},
								},
							},
							{
								ID:    5,
								Name:  "d",
								Links: []*lib.Node{},
							},
						},
					},
					{
						ID:   1,
						Name: "b",
						Links: []*lib.Node{
							{
								ID:    3,
								Name:  "c",
								Links: []*lib.Node{},
							},
						},
					},
					{
						ID:    3,
						Name:  "c",
						Links: []*lib.Node{},
					},
				},
			},
			IsError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got, err := JustCostRouteWithDijkstra(test.Input.graph, test.Input.startName, test.Input.targetCost, test.Input.getCost)
			if test.IsError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.Expected, got)
		})
	}
}
