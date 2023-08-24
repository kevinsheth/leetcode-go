package leetcode

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		k    int
		want []int
	}{
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},
	}

	for _, test := range tests {
		k := RemoveElement(test.nums, test.val)
		if k != test.k {
			t.Errorf("RemoveElement() returned length %v, want length %v", k, test.k)
		}

		for i := 0; i < test.k; i++ {
			if test.nums[i] != test.want[i] {
				t.Errorf("RemoveElement() modified array %v, want first %v elements to be %v", test.nums, test.k, test.want)
				break
			}
		}
	}
}
