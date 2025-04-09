package problem

import (
  "reflect"
  "testing"
)

/*
3375. Minimum Operations to Make Array Values Equal to K
Easy
Topics
Companies
Hint

You are given an integer array nums and an integer k.

An integer h is called valid if all values in the array that are strictly greater than h are identical.

For example, if nums = [10, 8, 10, 8], a valid integer is h = 9 because all nums[i] > 9 are equal to 10, but 5 is not a valid integer.

You are allowed to perform the following operation on nums:

    Select an integer h that is valid for the current values in nums.
    For each index i where nums[i] > h, set nums[i] to h.

Return the minimum number of operations required to make every element in nums equal to k. If it is impossible to make all elements equal to k, return -1.
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
      []int{5,2,5,4,5},
      2,
      2,
    },
    {
      "test2",
      []int{2,1,2},
      2,
      -1,
    },
    {
      "test3",
      []int{5,5,5,6,7,8,9},
      4,
      5,
    },
    {
      "test4",
      []int{5,5,5,6,7,8,9},
      6,
      -1,
    },

  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := minOperations(tc.nums, tc.k)
      ok := reflect.DeepEqual(got, tc.want)
      if !ok {
        t.Errorf("minOperations(%v, %v) = %v; want = %v\n",
        tc.nums, tc.k, got, tc.want)
      }
    })
  }
}
