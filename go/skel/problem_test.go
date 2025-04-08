package problem

import (
  "reflect"
  "testing"
)

/*
*/ 


func TestProblem(t *testing.T) {

  testCases := []struct {
    name string 
    nums []int
    want int
  } {
    {
      "test1",
      []int{1,2,3},
      0,
    },
    {
      "test2",
      []int{1,2,3},
      0,
    },

  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := solve(tc.nums)
      ok := reflect.DeepEqual(got, tc.want)
      if !ok {
        t.Errorf("solve(%v) = %v; want = %v\n",
        tc.nums, got, tc.want)
      }
    })
  }
}
