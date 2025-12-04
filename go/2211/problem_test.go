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
    directions string
    want int
  } {
    {
      "test1",
      "RLRSLL",
      5,
    },
    {
      "test2",
      "LLRR",
      0,
    },

  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := solve(tc.directions)
      ok := reflect.DeepEqual(got, tc.want)
      if !ok {
        t.Errorf("solve(%v) = %v; want = %v\n",
        tc.directions, got, tc.want)
      }
    })
  }
}
