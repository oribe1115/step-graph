package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitEdgeCost(t *testing.T) {
	tests := []struct {
		Label    string
		Input    bool
		Expected *EdgeCost
	}{
		{
			Label: "SUCCESS: directed",
			Input: true,
			Expected: &EdgeCost{
				Costs:      map[int]map[int]int{},
				IsDirected: true,
			},
		},
		{
			Label: "SUCCESS: undirected",
			Input: false,
			Expected: &EdgeCost{
				Costs:      map[int]map[int]int{},
				IsDirected: false,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := InitEdgeCost(test.Input)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestEdgeCostCheckIDOrder(t *testing.T) {
	type input struct {
		idA int
		idB int
	}
	type expected struct {
		firstID  int
		secondID int
	}
	tests := []struct {
		Label    string
		Use      *EdgeCost
		Input    input
		Expected expected
	}{
		{
			Label: "SUCCESS: directed & idA < idB",
			Use: &EdgeCost{
				IsDirected: true,
			},
			Input:    input{1, 2},
			Expected: expected{1, 2},
		},
		{
			Label: "SUCCESS: directed & idA > idB",
			Use: &EdgeCost{
				IsDirected: true,
			},
			Input:    input{2, 1},
			Expected: expected{2, 1},
		},
		{
			Label: "SUCCESS: undirected & idA < idB",
			Use: &EdgeCost{
				IsDirected: false,
			},
			Input:    input{1, 2},
			Expected: expected{1, 2},
		},
		{
			Label: "SUCCESS: undirected & idA > idB",
			Use: &EdgeCost{
				IsDirected: false,
			},
			Input:    input{2, 1},
			Expected: expected{1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			firstID, secondID := test.Use.checkIDOrder(test.Input.idA, test.Input.idB)
			assert.Equal(t, test.Expected.firstID, firstID)
			assert.Equal(t, test.Expected.secondID, secondID)
		})
	}
}

func TestEdgeCostSet(t *testing.T) {
	type input struct {
		firstID  int
		secondID int
		cost     int
	}
	tests := []struct {
		Label    string
		Use      *EdgeCost
		Input    input
		Expected *EdgeCost
	}{
		{
			Label: "SUCCESS: directed & firstID is unset & firstID < secondID",
			Use: &EdgeCost{
				Costs:      map[int]map[int]int{},
				IsDirected: true,
			},
			Input: input{1, 2, 5},
			Expected: &EdgeCost{
				Costs: map[int]map[int]int{
					1: {2: 5},
				},
				IsDirected: true,
			},
		},
		{
			Label: "SUCCESS: directed & firstID is unset & firstID > secondID",
			Use: &EdgeCost{
				Costs:      map[int]map[int]int{},
				IsDirected: true,
			},
			Input: input{2, 1, 5},
			Expected: &EdgeCost{
				Costs: map[int]map[int]int{
					2: {1: 5},
				},
				IsDirected: true,
			},
		},
		{
			Label: "SUCCESS: directed & firstID is set & firstID < secondID",
			Use: &EdgeCost{
				Costs: map[int]map[int]int{
					1: {3: 4},
				},
				IsDirected: true,
			},
			Input: input{1, 2, 5},
			Expected: &EdgeCost{
				Costs: map[int]map[int]int{
					1: {2: 5, 3: 4},
				},
				IsDirected: true,
			},
		},
		{
			Label: "SUCCESS: directed & firstID is set & firstID > secondID",
			Use: &EdgeCost{
				Costs: map[int]map[int]int{
					2: {3: 4},
				},
				IsDirected: true,
			},
			Input: input{2, 1, 5},
			Expected: &EdgeCost{
				Costs: map[int]map[int]int{
					2: {1: 5, 3: 4},
				},
				IsDirected: true,
			},
		},
		{
			Label: "SUCCESS: undirected & firstID is unset & firstID < secondID",
			Use: &EdgeCost{
				Costs:      map[int]map[int]int{},
				IsDirected: false,
			},
			Input: input{1, 2, 5},
			Expected: &EdgeCost{
				Costs: map[int]map[int]int{
					1: {2: 5},
				},
				IsDirected: false,
			},
		},
		{
			Label: "SUCCESS: undirected & firstID is unset & firstID > secondID",
			Use: &EdgeCost{
				Costs:      map[int]map[int]int{},
				IsDirected: false,
			},
			Input: input{2, 1, 5},
			Expected: &EdgeCost{
				Costs: map[int]map[int]int{
					1: {2: 5},
				},
				IsDirected: false,
			},
		},
		{
			Label: "SUCCESS: undirected & firstID is set & firstID < secondID",
			Use: &EdgeCost{
				Costs: map[int]map[int]int{
					1: {3: 4},
				},
				IsDirected: false,
			},
			Input: input{1, 2, 5},
			Expected: &EdgeCost{
				Costs: map[int]map[int]int{
					1: {2: 5, 3: 4},
				},
				IsDirected: false,
			},
		},
		{
			Label: "SUCCESS: undirected & firstID is set & firstID > secondID",
			Use: &EdgeCost{
				Costs: map[int]map[int]int{
					2: {3: 4},
				},
				IsDirected: false,
			},
			Input: input{2, 4, 5},
			Expected: &EdgeCost{
				Costs: map[int]map[int]int{
					2: {3: 4, 4: 5},
				},
				IsDirected: false,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			test.Use.Set(test.Input.firstID, test.Input.secondID, test.Input.cost)
			assert.Equal(t, test.Expected, test.Use)
		})
	}
}

func TestEdgeCostGet(t *testing.T) {
	type input struct {
		firstID  int
		secondID int
	}
	type expected struct {
		cost int
		ok   bool
	}
	tests := []struct {
		Label    string
		Use      *EdgeCost
		Input    input
		Expected expected
	}{
		{
			Label: "SUCCESS: set data",
			Use: &EdgeCost{
				Costs: map[int]map[int]int{
					1: {2: 5},
				},
				IsDirected: false,
			},
			Input:    input{1, 2},
			Expected: expected{5, true},
		},
		{
			Label: "FAIL: unset data",
			Use: &EdgeCost{
				Costs: map[int]map[int]int{
					1: {3: 4},
				},
				IsDirected: false,
			},
			Input:    input{1, 2},
			Expected: expected{0, false},
		},
		{
			Label: "FAIL: unset firstID",
			Use: &EdgeCost{
				Costs:      map[int]map[int]int{},
				IsDirected: false,
			},
			Input:    input{1, 2},
			Expected: expected{0, false},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			cost, ok := test.Use.Get(test.Input.firstID, test.Input.secondID)
			assert.Equal(t, test.Expected.cost, cost)
			assert.Equal(t, test.Expected.ok, ok)
		})
	}
}
