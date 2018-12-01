package chapter02_test

import (
	"CLRS/chapter02"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortDesc(t *testing.T) {
	functions := map[string]func([]int) ([]int, error){
		"InsertionSortDesc": chapter02.InsertionSortDesc,
		"BubbleSort":        chapter02.BubbleSort,
		"SelectionSort":     chapter02.SelectionSort,
		"MergeSort":         chapter02.MergeSort,
	}

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

	for fn, f := range functions {
		for tn, tc := range testcases {
			out, err := f(tc.input)
			assert.Equalf(t, tc.output, out, "%s %s Output check", fn, tn)
			if tc.expectedErr != nil {
				assert.Equalf(t, tc.expectedErr, err.Error(), "%s Error Check", tn)
			}
		}
	}
}

func TestSortAsc(t *testing.T) {
	functions := map[string]func([]int) ([]int, error){
		"InsertionSortAsc": chapter02.InsertionSortAsc,
	}
	testcases := map[string]struct {
		input       []int
		output      []int
		expectedErr error
	}{
		"Already Sorted": {
			input:  []int{3, 2, 1},
			output: []int{3, 2, 1},
		},
		"No Error": {
			input:  []int{1, 2, 3},
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

	for fn, f := range functions {
		for tn, tc := range testcases {
			out, err := f(tc.input)
			assert.Equalf(t, tc.output, out, "%s %s Output check", fn, tn)
			if tc.expectedErr != nil {
				assert.Equalf(t, tc.expectedErr, err.Error(), "%s Error Check", tn)
			}
		}
	}
}
