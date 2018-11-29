package chapter02_test

import (
	"CLRS/chapter02"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSortDesc(t *testing.T) {
	testcases := map[string]struct {
		input       []int
		output      []int
		expectedErr error
	}{
		"No Error": {
			input:  []int{3, 2, 1},
			output: []int{1, 2, 3},
		},
		"Already Sorted": {
			input:  []int{1, 2, 3},
			output: []int{1, 2, 3},
		},
		"Double Up": {
			input:  []int{3, 1, 1, 2},
			output: []int{1, 1, 2, 3},
		},
		"Book example": {
			input:  []int{31, 41, 59, 26, 41, 58},
			output: []int{26, 31, 41, 41, 58, 59},
		},
		"Empty": {
			input:  []int{},
			output: []int{},
		},
	}

	for tn, tc := range testcases {
		out, err := chapter02.InsertionSortDesc(tc.input)
		assert.Equal(t, tc.output, out)
		if tc.expectedErr != nil {
			assert.Equalf(t, tc.expectedErr, err.Error(), "%s Error Check", tn)
		}
	}
}

func TestInsertionSortAsc(t *testing.T) {
	testcases := map[string]struct {
		input       []int
		output      []int
		expectedErr error
	}{
		"No Error": {
			input:  []int{1, 2, 3},
			output: []int{3, 2, 1},
		},
		"Already Sorted": {
			input:  []int{3, 2, 1},
			output: []int{3, 2, 1},
		},
		"Double Up": {
			input:  []int{3, 1, 1, 2},
			output: []int{3, 2, 1, 1},
		},
		"Book example": {
			input:  []int{31, 41, 59, 26, 41, 58},
			output: []int{59, 58, 41, 41, 31, 26},
		},
		"Empty": {
			input:  []int{},
			output: []int{},
		},
	}

	for tn, tc := range testcases {
		out, err := chapter02.InsertionSortAsc(tc.input)
		assert.Equal(t, tc.output, out)
		if tc.expectedErr != nil {
			assert.Equalf(t, tc.expectedErr, err.Error(), "%s Error Check", tn)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	testcases := map[string]struct {
		input       []int
		output      []int
		expectedErr error
	}{
		"No Error": {
			input:  []int{3, 2, 1},
			output: []int{1, 2, 3},
		},
		"Already Sorted": {
			input:  []int{1, 2, 3},
			output: []int{1, 2, 3},
		},
		"Double Up": {
			input:  []int{3, 1, 1, 2},
			output: []int{1, 1, 2, 3},
		},
		"Book example": {
			input:  []int{31, 41, 59, 26, 41, 58},
			output: []int{26, 31, 41, 41, 58, 59},
		},
		"Empty": {
			input:  []int{},
			output: []int{},
		},
	}

	for tn, tc := range testcases {
		out, err := chapter02.SelectionSort(tc.input)
		assert.Equal(t, tc.output, out)
		if tc.expectedErr != nil {
			assert.Equalf(t, tc.expectedErr, err.Error(), "%s Error Check", tn)
		}
	}
}
