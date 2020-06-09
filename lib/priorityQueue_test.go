package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	type item struct {
		node     *Node
		from     *Node
		priority int
	}
	tests := []struct {
		Label    string
		Input    []item
		Expected []item
	}{
		{
			Label: "SUCCESS: normal",
			Input: []item{
				{
					node:     &Node{ID: 1, Name: "a", Links: []*Node{}},
					from:     nil,
					priority: 10,
				},
				{
					node:     &Node{ID: 2, Name: "b", Links: []*Node{}},
					from:     &Node{ID: 1, Name: "a", Links: []*Node{}},
					priority: 1,
				},
				{
					node:     &Node{ID: 3, Name: "c", Links: []*Node{}},
					from:     &Node{ID: 2, Name: "b", Links: []*Node{}},
					priority: 5,
				},
			},
			Expected: []item{
				{
					node:     &Node{ID: 2, Name: "b", Links: []*Node{}},
					from:     &Node{ID: 1, Name: "a", Links: []*Node{}},
					priority: 1,
				},
				{
					node:     &Node{ID: 3, Name: "c", Links: []*Node{}},
					from:     &Node{ID: 2, Name: "b", Links: []*Node{}},
					priority: 5,
				},
				{
					node:     &Node{ID: 1, Name: "a", Links: []*Node{}},
					from:     nil,
					priority: 10,
				},
			},
		},
		{
			Label: "Fial: pop is more than push",
			Input: []item{
				{
					node:     &Node{ID: 1, Name: "a", Links: []*Node{}},
					from:     nil,
					priority: 10,
				},
			},
			Expected: []item{
				{
					node:     &Node{ID: 1, Name: "a", Links: []*Node{}},
					from:     nil,
					priority: 10,
				},
				{
					node:     nil,
					from:     nil,
					priority: 0,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			pq := InitPriorityQueue()
			for _, input := range test.Input {
				pq.Push(input.node, input.from, input.priority)
			}

			for _, expected := range test.Expected {
				node, from, priority := pq.Pop()
				assert.Equal(t, expected.node, node)
				assert.Equal(t, expected.from, from)
				assert.Equal(t, expected.priority, priority)
			}
		})
	}
}
