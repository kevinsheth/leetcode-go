package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
		{
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
	}

	for _, test := range tests {
		got := MergeSortedArrays(test.nums1, test.m, test.nums2, test.n)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("MergeSortedArrays(%v, %d, %v, %d) = %v; want %v", test.nums1, test.m, test.nums2, test.n, got, test.want)

		}
	}

}
