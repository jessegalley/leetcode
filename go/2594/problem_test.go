package  problem 

import (
  "testing"
) 

func TestMechanics(t *testing.T) {
  testCases := []struct {
    name string 
    ranks []int 
    cars int
    want int64
  }{
    {
      "test1",
      []int{4,3,2,1},
      10,
      16,
    },
    {
      "test2",
      []int{5,1,8},
      6,
      16,
    },
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := repairCars(tc.ranks, tc.cars)
      if got != tc.want {
        t.Errorf("repairCars(%v, %d) = %d; want: %d",
          tc.ranks, tc.cars, got, tc.want)
      }
    })
  }
}

