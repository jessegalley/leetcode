package main

import (
  "fmt"
)

func twoSum(nums []int, target int) []int {
  numMap := make(map[int]int)
  for i, v := range nums {
    c := target - v 
    found, ok := numMap[c]
    if ok {
      return []int{found, i}
    }
    numMap[v] = i
   }

  return nil
}


func main () {

  nums := []int{2,7,11,5}
  target := 9


  result := twoSum(nums, target)

  fmt.Println("result: ", result)
  


}
