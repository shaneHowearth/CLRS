package chapter02_test

import (
	"CLRS/chapter02"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {
	testcases := map[string]struct {
		find        int
		input       []int
		output      int
		expectedErr error
	}{
		"No Error": {
			find:   2,
			input:  []int{3, 2, 1},
			output: 1,
		},
		"Sorted": {
			find:   2,
			input:  []int{1, 2, 3},
			output: 1,
		},
		"Nothing to find": {
			find:   2,
			input:  []int{},
			output: -1,
		},
		"Non existant": {
			find:   2,
			input:  []int{31, 41, 59, 26, 41, 58},
			output: -1,
		},
	}

	for tn, tc := range testcases {
		out, err := chapter02.LinearSearch(tc.find, tc.input)
		assert.Equal(t, tc.output, out)
		if tc.expectedErr != nil {
			assert.Equalf(t, tc.expectedErr, err.Error(), "%s Error Check", tn)
		} else {
			assert.NoErrorf(t, err, "%s Error Check", tn)
		}
	}
}
