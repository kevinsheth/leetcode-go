package leetcode;

func RemoveElement(nums []int, val int) int {

  l, r := 0, 0

  for r < len(nums) {

    if nums[r] != val {
      nums[l] = nums[r]
      l++
    }
    r++
  }

  return l

}
