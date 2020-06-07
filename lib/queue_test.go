package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitQueue(t *testing.T) {
	tests := []struct {
		Label    string
		Expected *Queue
	}{
		{
			Label: "SUCCESS",
			Expected: &Queue{
				Nodes: make([]*Node, 0),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := InitQueue()
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestQueueLen(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *Queue
		Expected int
	}{
		{
			Label: "SUCCESS: normal",
			Use: &Queue{
				Nodes: []*Node{
					{
						1,
						"a",
						[]*Node{},
					},
					{
						2,
						"b",
						[]*Node{},
					},
				},
			},
			Expected: 2,
		},
		{
			Label: "SUCCESS: zero",
			Use: &Queue{
				Nodes: []*Node{},
			},
			Expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := test.Use.Len()
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestQueueSprint(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *Queue
		Expected string
	}{
		{
			Label: "SUCCESS: normal",
			Use: &Queue{
				Nodes: []*Node{
					{
						1,
						"a",
						[]*Node{},
					},
					{
						2,
						"b",
						[]*Node{},
					},
				},
			},
			Expected: "{1: a} {2: b} ",
		},
		{
			Label: "SUCCESS: length is zero",
			Use: &Queue{
				Nodes: []*Node{},
			},
			Expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := test.Use.Sprint()
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestQueueEnqueue(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *Queue
		Input    *Node
		Expected *Queue
	}{
		{
			Label: "SUCCESS: normal",
			Use: &Queue{
				Nodes: []*Node{
					{
						1,
						"a",
						[]*Node{},
					},
					{
						2,
						"b",
						[]*Node{},
					},
				},
			},
			Input: &Node{
				3,
				"c",
				[]*Node{},
			},
			Expected: &Queue{
				Nodes: []*Node{
					{
						1,
						"a",
						[]*Node{},
					},
					{
						2,
						"b",
						[]*Node{},
					},
					{
						3,
						"c",
						[]*Node{},
					},
				},
			},
		},
		{
			Label: "SUCCESS: first node",
			Use: &Queue{
				Nodes: []*Node{},
			},
			Input: &Node{
				1,
				"a",
				[]*Node{},
			},
			Expected: &Queue{
				Nodes: []*Node{
					{
						1,
						"a",
						[]*Node{},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			test.Use.Enqueue(test.Input)
			assert.Equal(t, test.Expected, test.Use)
		})
	}
}

func TestQueueDequeue(t *testing.T) {
	tests := []struct {
		Label         string
		Use           *Queue
		ExpectedGot   *Node
		ExpectedAfter *Queue
		IsError       bool
	}{
		{
			Label: "SUCCESS: normal",
			Use: &Queue{
				Nodes: []*Node{
					{
						1,
						"a",
						[]*Node{},
					},
					{
						2,
						"b",
						[]*Node{},
					},
				},
			},
			ExpectedGot: &Node{
				1,
				"a",
				[]*Node{},
			},
			ExpectedAfter: &Queue{
				Nodes: []*Node{
					{
						2,
						"b",
						[]*Node{},
					},
				},
			},
			IsError: false,
		},
		{
			Label: "SUCCESSS: last node",
			Use: &Queue{
				Nodes: []*Node{
					{
						1,
						"a",
						[]*Node{},
					},
				},
			},
			ExpectedGot: &Node{
				1,
				"a",
				[]*Node{},
			},
			ExpectedAfter: &Queue{
				Nodes: []*Node{},
			},
			IsError: false,
		},
		{
			Label: "FAIL: zero length",
			Use: &Queue{
				Nodes: []*Node{},
			},
			ExpectedGot: nil,
			ExpectedAfter: &Queue{
				Nodes: []*Node{},
			},
			IsError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got, err := test.Use.Dequeue()
			if test.IsError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, test.ExpectedGot, got)
			assert.Equal(t, test.ExpectedAfter, test.Use)
		})
	}
}
