package problem

import (
  "reflect"
  "testing"
)

/*
2176. Count Equal and Divisible Pairs in an Array
Easy
Topics
Companies
Hint
Given a 0-indexed integer array nums of length n and an integer k, return the number of pairs (i, j) where 0 <= i < j < n, such that nums[i] == nums[j] and (i * j) is divisible by k.

*/ 


func TestProblem(t *testing.T) {

  testCases := []struct {
    name string 
    nums []int
    k int 
    want int
  } {
    {
      "test1",
      []int{3,1,2,2,2,1,3},
      2,
      4,
    },
    {
      "test2",
      []int{1,2,3,4},
      1,
      0,
    },
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := countPairs(tc.nums, tc.k)
      ok := reflect.DeepEqual(got, tc.want)
      if !ok {
        t.Errorf("solve(%v, %d) = %v; want = %v\n",
         tc.nums, tc.k, got, tc.want)
      }
    })
  }
}
