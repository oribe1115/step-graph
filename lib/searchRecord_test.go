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
			Label:    "SUCCESS: normal",
			Expected: &SearchRecord{},
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
					1: {ID: 1, Count: 1},
					2: {ID: 2, Count: 2},
				},
			},
			Input:    1,
			Expected: true,
		},
		{
			Label: "SUCCESS: not found",
			Use: &SearchRecord{
				Records: map[int]*Record{
					1: {ID: 1, Count: 1},
					2: {ID: 2, Count: 2},
				},
			},
			Input:    3,
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
		ID    int
		Count int
	}
	tests := []struct {
		Label    string
		Use      *SearchRecord
		Input    input
		Expected *SearchRecord
		IsError  bool
	}{
		{
			Label: "SUCCESS: normal",
			Use: &SearchRecord{
				Records: map[int]*Record{
					1: {ID: 1, Count: 1},
					2: {ID: 2, Count: 2},
				},
			},
			Input: input{ID: 3, Count: 3},
			Expected: &SearchRecord{
				Records: map[int]*Record{
					1: {ID: 1, Count: 1},
					2: {ID: 2, Count: 2},
					3: {ID: 3, Count: 3},
				},
			},
			IsError: false,
		},
		{
			Label: "FAIL: deplicate id",
			Use: &SearchRecord{
				Records: map[int]*Record{
					1: {ID: 1, Count: 1},
					2: {ID: 2, Count: 2},
				},
			},
			Input:    input{ID: 1, Count: 3},
			Expected: nil,
			IsError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			err := test.Use.AddRecord(test.Input.ID, test.Input.Count)
			if test.IsError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
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
					1: {ID: 1, Count: 1},
					2: {ID: 2, Count: 2},
				},
			},
			Input: 1,
			Expected: &Record{
				ID:    1,
				Count: 1,
			},
		},
		{
			Label: "FAIL: not found",
			Use: &SearchRecord{
				Records: map[int]*Record{
					1: {ID: 1, Count: 1},
					2: {ID: 2, Count: 2},
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
