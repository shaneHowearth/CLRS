package chapter02_test

import (
	"CLRS/chapter02"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
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
	}

	for tn, tc := range testcases {
		out, err := chapter02.InsertionSort(tc.input)
		assert.Equal(t, tc.output, out)
		if tc.expectedErr != nil {
			assert.Equalf(t, tc.expectedErr, err.Error(), "%s Error Check", tn)
		}
	}
}
