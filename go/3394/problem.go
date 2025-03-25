package problem

import (
	"cmp"
	"slices"
)


func checkValidCuts(n int, rectangles [][]int) bool {
  
  var cx [][]int
  var cy [][]int

  // every start coord will be worth 1
  // every end coord wil be worth -1 
  for _, rect := range rectangles {
    cx = append(cx, []int{rect[0], 1})
    cx = append(cx, []int{rect[2], -1})
    cy = append(cy, []int{rect[1], 1})
    cy = append(cy, []int{rect[3], -1})
  }

  // define a sort func for position then val  
  mySortFunc := func(a, b []int) int {
    // first compare by [0]
    if firstCmp := cmp.Compare(a[0], b[0]); firstCmp != 0 {
      return firstCmp
    }

    // if firstCmp is 0, the values are the same 
    // so we return the comparison of [1]
    return cmp.Compare(a[1], b[1])
  }

  // sort the two coord slices 
  slices.SortFunc(cx, mySortFunc)
  slices.SortFunc(cy, mySortFunc)

  return (countOverlaps(cx) >= 3 || countOverlaps(cy) >= 3) 
}

func countOverlaps(coords [][]int) int {
  lineSlices := 0
  overlapCnt := 0

  for _, coord := range coords {
    overlapCnt += coord[1]
    if overlapCnt == 0 {
      lineSlices++
    }
  }

  return lineSlices
}
