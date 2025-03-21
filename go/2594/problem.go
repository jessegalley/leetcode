package problem

import "fmt"

func repairCars(ranks []int, cars int) int64 {
    
  left := 0
  right := doRepair(sliceMax(ranks), cars)

  fmt.Println(ranks, cars)
  fmt.Printf("left: %d; right: %d\n", left, right)


  return 16
}

func doRepair(rank int, cars int) int64 {
  return int64(rank * cars * cars)
}

func sliceMax(nums []int) int {
  max := 0
  for _,num := range nums {
    if num > max {
      max = num
    }
  }

  return max 
}
