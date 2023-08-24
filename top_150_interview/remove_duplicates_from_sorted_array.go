package leetcode

func RemoveDuplicatesFromSortedArray(nums []int) int {

	l, r := 1, 1

	for r < len(nums) {

		if nums[r] != nums[r-1] {
			nums[l] = nums[r]
			l++
		}
		r++

	}
	return l

}
