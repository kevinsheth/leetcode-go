package leetcode

func MergeSortedArrays(nums1 []int, m int, nums2 []int, n int) []int {

	p1 := m - 1
	p2 := n - 1

	for p := len(nums1) - 1; p >= 0; p-- {

		if p2 < 0 {
			break
		}

		if p1 >= 0 && nums1[p1] >= nums2[p2] {
			nums1[p] = nums1[p1]
			p1 -= 1
		} else {
			nums1[p] = nums2[p2]
			p2 -= 1
		}
	}

	return nums1
}
