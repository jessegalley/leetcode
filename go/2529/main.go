// leetcode problem #2529
package main

import (
  "fmt"
)


func upperBound(ints []int) int {
  start := 0
  end := len(ints) -1 
  i := end +1 

  for {
    // fmt.Printf("start: %d, i: %d, end: %d\n", start, i , end)
    if start > end {
      break
    }

    mid := (start + end) / 2 

    if ints[mid] <= 0 {
      start = mid + 1
    } else if ints[mid] > 0 {
      end = mid - 1
      i = mid
    }
  }

  return i
}

func lowerBound(ints []int) int {
  start := 0
  end := len(ints) -1 
  i := end +1

  for {
    if start > end {
      break
    }

    mid := (start + end) / 2 

    if ints[mid] < 0 {
      start = mid + 1
    } else if ints[mid] >= 0 {
      end = mid - 1
      i = mid
    }
  }

  return i
}

func maximumCount(nums []int) int {
  upper := upperBound(nums)
  lower := lowerBound(nums)

  // fmt.Printf("u: %d, l: %d\n", upper, lower)

  cntPos := len(nums) - upper
  cntNeg := lower    

  // fmt.Printf("p: %d, d: %d\n", cntPos, cntNeg)

  return max(cntPos, cntNeg)
}

func main() {
  // var testcase = []int{-2,-2,-1,1,2,2,3,4,5,6}
  var testcase = []int{0,0,0,0,0,0,0,0}

  max := maximumCount(testcase)

  fmt.Println("max: ", max)

  
}



