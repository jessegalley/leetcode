package problem

import (
  "reflect"
  "testing"
)

/*
3396. Minimum Number of Operations to Make Elements in Array Distinct
Easy
Topics
Companies
Hint

You are given an integer array nums. You need to ensure that the elements in the array are distinct. To achieve this, you can perform the following operation any number of times:

    Remove 3 elements from the beginning of the array. If the array has fewer than 3 elements, remove all remaining elements.

Note that an empty array is considered to have distinct elements. Return the minimum number of operations needed to make the elements in the array distinct.
*/ 


func TestProblem(t *testing.T) {

  testCases := []struct {
    name string 
    nums []int
    want int
  } {
    {
      "test1",
      []int{1,2,3,4,2,3,3,5,7},
      2,
    },
    {
      "test2",
      []int{4,5,6,4,4},
      2,
    },
    {
      "test3",
      []int{6,7,8,9},
      0,
    },
    {
      "test4",
      []int{5,5},
      1,
    },
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := minimumOperations(tc.nums)
      ok := reflect.DeepEqual(got, tc.want)
      if !ok {
        t.Errorf("minimumOperations(%v) = %v; want = %v\n",
        tc.nums, got, tc.want)
      }
    })
  }
}
