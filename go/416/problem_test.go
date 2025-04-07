package problem

import (
	"reflect"
	"testing"
)

/*
416. Partition Equal Subset Sum
Medium
Topics
Companies

Given an integer array nums, return true if you can partition the array into two subsets such that the sum of the elements in both subsets is equal or false otherwise.
*/ 


func TestProblem(t *testing.T) {

  testCases := []struct {
    name string 
    nums []int 
    want bool
  } {
    {
      "test1",
      []int{1,5,11,5},
      true,
    },
    {
      "test2",
      []int{1,2,3,5},
      false,
    },
    {
      "test3",
      []int{10},
      false,
    },

  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := canPartition(tc.nums)
      ok := reflect.DeepEqual(got, tc.want)
      if !ok {
        t.Errorf("canPartition(%v) = %v; want = %v\n",
          tc.nums, got, tc.want)
      }
    })
  }
}
