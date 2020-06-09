package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitSearchRecord(t *testing.T) {
	tests := []struct {
		Label    string
		Expected *SearchRecord
	}{
		{
			Label: "SUCCESS: normal",
			Expected: &SearchRecord{
				Records: map[int]*Record{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := InitSearchRecord()
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestSearchRecordIsRecorded(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *SearchRecord
		Input    int
		Expected bool
	}{
		{
			Label: "SUCCESS: found",
			Use: &SearchRecord{
				Records: map[int]*Record{
					1: {
						Node: &Node{
							ID:    1,
							Name:  "a",
							Links: []*Node{},
						},
						Count: 1,
						From:  nil,
					},
					2: {
						Node: &Node{
							ID:    2,
							Name:  "b",
							Links: []*Node{},
						},
					},
				},
			},
			Input:    1,
			Expected: true,
		},
		{
			Label: "SUCCESS: not found",
			Use: &SearchRecord{
				Records: map[int]*Record{
					1: {
						Node: &Node{
							ID:    1,
							Name:  "a",
							Links: []*Node{},
						},
						Count: 1,
						From:  nil,
					},
					2: {
						Node: &Node{
							ID:    2,
							Name:  "b",
							Links: []*Node{},
						},
						Count: 2,
						From:  nil,
					},
				},
			},
			Input:    3,
			Expected: false,
		},
		{
			Label: "SUCCESS: no rocord",
			Use: &SearchRecord{
				Records: map[int]*Record{},
			},
			Input:    2,
			Expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := test.Use.IsRecorded(test.Input)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestAddRecord(t *testing.T) {
	type input struct {
		node  *Node
		count int
		from  *Node
	}
	tests := []struct {
		Label    string
		Use      *SearchRecord
		Input    input
		Expected *SearchRecord
	}{
		{
			Label: "SUCCESS: normal",
			Use: &SearchRecord{
				Records: map[int]*Record{
					1: {
						Node: &Node{
							ID:    1,
							Name:  "a",
							Links: []*Node{},
						},
						Count: 1,
						From:  nil,
					},
				},
			},
			Input: input{
				node: &Node{
					ID:    2,
					Name:  "b",
					Links: []*Node{},
				},
				count: 2,
				from: &Node{
					ID:    1,
					Name:  "a",
					Links: []*Node{},
				},
			},
			Expected: &SearchRecord{
				Records: map[int]*Record{
					1: {
						Node: &Node{
							ID:    1,
							Name:  "a",
							Links: []*Node{},
						},
						Count: 1,
						From:  nil,
					},
					2: {
						Node: &Node{
							ID:    2,
							Name:  "b",
							Links: []*Node{},
						},
						Count: 2,
						From: &Node{
							ID:    1,
							Name:  "a",
							Links: []*Node{},
						},
					},
				},
			},
		},
		{
			Label: "SUCCESS: from is nil",
			Use: &SearchRecord{
				Records: map[int]*Record{
					1: {
						Node: &Node{
							ID:    1,
							Name:  "a",
							Links: []*Node{},
						},
						Count: 1,
						From:  nil,
					},
				},
			},
			Input: input{
				node: &Node{
					ID:    2,
					Name:  "b",
					Links: []*Node{},
				},
				count: 2,
				from:  nil,
			},
			Expected: &SearchRecord{
				Records: map[int]*Record{
					1: {
						Node: &Node{
							ID:    1,
							Name:  "a",
							Links: []*Node{},
						},
						Count: 1,
						From:  nil,
					},
					2: {
						Node: &Node{
							ID:    2,
							Name:  "b",
							Links: []*Node{},
						},
						Count: 2,
						From:  nil,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			test.Use.AddRecord(test.Input.node, test.Input.count, test.Input.from)
			assert.Equal(t, test.Expected, test.Use)
		})
	}
}

func TestGetRecord(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *SearchRecord
		Input    int
		Expected *Record
	}{
		{
			Label: "SUCCESS: normal",
			Use: &SearchRecord{
				Records: map[int]*Record{
					1: {
						Node: &Node{
							ID:    1,
							Name:  "a",
							Links: []*Node{},
						},
						Count: 1,
						From:  nil,
					},
				},
			},
			Input: 1,
			Expected: &Record{
				Node: &Node{
					ID:    1,
					Name:  "a",
					Links: []*Node{},
				},
				Count: 1,
				From:  nil,
			},
		},
		{
			Label: "FAIL: not found",
			Use: &SearchRecord{
				Records: map[int]*Record{
					1: {
						Node: &Node{
							ID:    1,
							Name:  "a",
							Links: []*Node{},
						},
						Count: 1,
						From:  nil,
					},
				},
			},
			Input:    3,
			Expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := test.Use.GetRecord(test.Input)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestSearchRecordGetRoute(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *SearchRecord
		Input    *Node
		Expected []*Node
		IsError  bool
	}{
		{
			Label: "SUCCESS: normal",
			Use: &SearchRecord{
				Records: map[int]*Record{
					1: {
						Node: &Node{
							ID:    1,
							Name:  "a",
							Links: []*Node{},
						},
						Count: 1,
						From:  nil,
					},
					2: {
						Node: &Node{
							ID:    2,
							Name:  "b",
							Links: []*Node{},
						},
						Count: 2,
						From: &Node{
							ID:    1,
							Name:  "a",
							Links: []*Node{},
						},
					},
					3: {
						Node: &Node{
							ID:    3,
							Name:  "c",
							Links: []*Node{},
						},
						Count: 3,
						From: &Node{
							ID:    2,
							Name:  "b",
							Links: []*Node{},
						},
					},
				},
			},
			Input: &Node{
				ID:    3,
				Name:  "c",
				Links: []*Node{},
			},
			Expected: []*Node{
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
				{
					ID:    3,
					Name:  "c",
					Links: []*Node{},
				},
			},
			IsError: false,
		},
		{
			Label: "FAIL: find infinit loop",
			Use: &SearchRecord{
				Records: map[int]*Record{
					1: {
						Node: &Node{
							ID:    1,
							Name:  "a",
							Links: []*Node{},
						},
						Count: 1,
						From: &Node{
							ID:    2,
							Name:  "b",
							Links: []*Node{},
						},
					},
					2: {
						Node: &Node{
							ID:    2,
							Name:  "b",
							Links: []*Node{},
						},
						Count: 2,
						From: &Node{
							ID:    1,
							Name:  "a",
							Links: []*Node{},
						},
					},
					3: {
						Node: &Node{
							ID:    3,
							Name:  "c",
							Links: []*Node{},
						},
						Count: 3,
						From: &Node{
							ID:    2,
							Name:  "b",
							Links: []*Node{},
						},
					},
				},
			},
			Input: &Node{
				ID:    3,
				Name:  "c",
				Links: []*Node{},
			},
			Expected: nil,
			IsError:  true,
		},
		{
			Label: "FAIL: faild to find record",
			Use: &SearchRecord{
				Records: map[int]*Record{
					1: {
						Node: &Node{
							ID:    1,
							Name:  "a",
							Links: []*Node{},
						},
						Count: 1,
						From:  nil,
					},
					2: {
						Node: &Node{
							ID:    2,
							Name:  "b",
							Links: []*Node{},
						},
						Count: 2,
						From: &Node{
							ID:    4,
							Name:  "d",
							Links: []*Node{},
						},
					},
					3: {
						Node: &Node{
							ID:    3,
							Name:  "c",
							Links: []*Node{},
						},
						Count: 3,
						From: &Node{
							ID:    2,
							Name:  "b",
							Links: []*Node{},
						},
					},
				},
			},
			Input: &Node{
				ID:    3,
				Name:  "c",
				Links: []*Node{},
			},
			Expected: nil,
			IsError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got, err := test.Use.GetRoute(test.Input)
			if test.IsError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.Expected, got)
		})
	}
}
