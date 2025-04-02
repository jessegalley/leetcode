package problem

import (
	"reflect"
	"testing"
)

/*
2873. Maximum Value of an Ordered Triplet I
Easy
Topics
Companies
Hint

You are given a 0-indexed integer array nums.

Return the maximum value over all triplets of indices (i, j, k) such that i < j < k. If all such triplets have a negative value, return 0.

The value of a triplet of indices (i, j, k) is equal to (nums[i] - nums[j]) * nums[k].
*/ 


func TestProblem(t *testing.T) {

  testCases := []struct {
    name string 
    nums []int 
    want int64 
  } {
    {
      "test1",
      []int{12,6,1,2,7},
      77,
    },
    {
      "test2",
      []int{1,10,3,4,19},
      133,
    },
    {
      "test3",
      []int{1,2,3},
      0,
    },
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := maximumTripletValue(tc.nums)
      ok := reflect.DeepEqual(got, tc.want)
      if !ok {
        t.Errorf("maximumTripletValue(%v) = %v; want = %v\n",
          tc.nums, got, tc.want)
      }
    })
  }
}
