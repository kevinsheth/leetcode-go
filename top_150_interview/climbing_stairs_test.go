package leetcode

import (
  "testing"
)

func TestClimbingStairs(t *testing.T) {
  tests := []struct {
    n int
    w int
  }{
    {2, 2},
    {3, 3},
  }

  for _, test := range tests {
    if climbingStairs(test.n) != test.w {
      t.Errorf("climbingStairs(%d) = %d; want %d", test.n, climbingStairs(test.n), test.w)
    }
  }
}
