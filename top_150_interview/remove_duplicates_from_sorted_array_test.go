package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
		output   []int
	}{
		{
			input:    []int{1, 1, 2},
			expected: 2,
			output:   []int{1, 2},
		},
		{
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
			output:   []int{0, 1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		nums := make([]int, len(test.input))
		copy(nums, test.input)
		k := RemoveDuplicatesFromSortedArray(nums)
		if k != test.expected || !reflect.DeepEqual(nums[:k], test.output) {
			t.Errorf("removeDuplicatesFromSortedArray(%v) = %d, %v; want %d, %v", test.input, k, nums[:k], test.expected, test.output)
		}
	}
}
