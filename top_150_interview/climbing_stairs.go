package leetcode

func climbingStairs(n int) int {

    l, r := 1, 1

    for i := 0; i < n - 1 ; i ++ {
        l, r = l + r, l
    }

    return l
}
