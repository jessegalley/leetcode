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
    input string 
    want int
  } {
    {
      "test1",
      "abcabcbb",
      3,
    },
    {
      "test2",
      "bbbb",
      1,
    },
    {
      "test2",
      "pwwkew",
      3,
    },

  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := solve(tc.input)
      ok := reflect.DeepEqual(got, tc.want)
      if !ok {
        t.Errorf("solve(%v) = %v; want = %v\n",
        tc.input, got, tc.want)
      }
    })
  }
}
