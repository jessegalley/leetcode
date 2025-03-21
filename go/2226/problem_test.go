package problem

import "testing"

func TestResult(t *testing.T) {
  
  testCases := []struct{
    name    string
    candies []int 
    k       int64
    wants   int
  }{
    {
      name: "example 1",
      candies: []int{5,8,6},
      k: 3,
      wants: 5,
    },
    {
      name: "example 2",
      candies: []int{2, 5},
      k: 11,
      wants: 0,
    },
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := maximumCandies(tc.candies, tc.k)
      if got != tc.wants {
        t.Errorf("maximumCandies(%v, %d) = %d; want %d",
          tc.candies, tc.k, got, tc.wants)
      }
    })
  }
}
